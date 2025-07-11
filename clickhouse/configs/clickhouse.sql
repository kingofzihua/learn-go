-- 链接 mysql 表
create table orders_source
(
    id            Int64,
    order_sn      Int64,
    user_id       Nullable(Int64),
    order_amount  Decimal(10, 2) comment '订单金额',
    is_refunded   Nullable(Bool) comment '是否退款, 0或1 (Bool类型更适合)',
    merchant_id   Int64 comment '商户ID',
    merchant_name Nullable(String) comment '商家名称',
    source        String comment '订单来源',
    payment_at    Nullable(DateTime) comment '支付时间',
    finish_at     Nullable(DateTime) comment '完成时间',
    order_status  Nullable(UInt8) comment '订单状态',
    created_at    DateTime default now(),
    updated_at    DateTime
)
    engine = MySQL('mariadb', 'learn_ck', 'orders', 'root', '[HIDDEN]');


-- 这个表用来存储每日按来源聚合的最终结果
CREATE TABLE order_source_daily_metrics
(
    -- 维度 (Dimensions)
    `stat_date` Date COMMENT '统计日期',
    `first_source` String COMMENT '订单的第一个来源',
    `source_path` String COMMENT '完整的来源路径',

    -- 指标 (Metrics)，使用 AggregateFunction 类型
    `order_count` AggregateFunction(count) COMMENT '订单数',
    `total_amount` AggregateFunction(sum, Decimal(10, 2)) COMMENT '订单总金额'

)
ENGINE = AggregatingMergeTree()
PARTITION BY toYYYYMM(stat_date)
ORDER BY (stat_date, first_source, source_path);


-- 创建物化视图，用于处理新数据并填充聚合表
CREATE MATERIALIZED VIEW order_source_mv TO order_source_daily_metrics
AS
SELECT
    -- 1. 聚合的粒度键：天
    created_at AS stat_date,

    -- 2. 字段转换逻辑
    -- first_source: 'banner:10,detail:200' -> 'banner'
    splitByChar(':', arrayElement(splitByChar(',', source), 1))[1] AS first_source,

    -- source_path: 'banner:10,detail:200,cart' -> 'banner,detail,cart'
    arrayStringConcat(
            arrayMap(
                    x -> splitByChar(':', x)[1],
                    splitByChar(',', source)
            ),
            ','
    ) AS source_path,

    -- 3. 聚合函数
    -- 使用 -State 后缀的聚合函数来生成中间状态
    countState() AS order_count,
    sumState(order_amount) AS total_amount

FROM orders_source
GROUP BY
    stat_date,
    first_source,
    source_path;

SELECT
    id,
    source,

    -- First Source
    splitByChar(':', arrayElement(splitByChar(',', source), 1))[1] AS first_source,

    -- Source Path
    arrayStringConcat(
            arrayMap(x -> splitByChar(':', x)[1], splitByChar(',', source)),
            ','
    ) AS source_path,

    -- Source Map (最复杂的部分)
    mapFromArrays(
            arrayMap(x -> x[1], arrayFilter(arr -> length(arr) = 2, arrayMap(s -> splitByChar(':', s), splitByChar(',', source)))),
            arrayMap(x -> x[2], arrayFilter(arr -> length(arr) = 2, arrayMap(s -> splitByChar(':', s), splitByChar(',', source))))
    ) AS source_map

FROM orders_source;

SELECT
    stat_date,
    first_source,
    source_path,
    countMerge(order_count) AS final_order_count,
    sumMerge(total_amount) AS final_total_amount
FROM order_source_daily_metrics
GROUP BY -- 别忘了 GROUP BY
         stat_date,
         first_source,
         source_path
ORDER BY
    stat_date DESC;

INSERT INTO order_source_daily_metrics
SELECT
    created_at AS stat_date,
    splitByChar(':', arrayElement(splitByChar(',', source), 1))[1] AS first_source,
    arrayStringConcat(
            arrayMap(
                    x -> splitByChar(':', x)[1],
                    splitByChar(',', source)
            ),
            ','
    ) AS source_path,
    countState() AS order_count,
    sumState(order_amount) AS total_amount
FROM orders_source
GROUP BY
    stat_date,
    first_source,
    source_path;


SELECT
    source_path,
    countMerge(order_count) AS totals
FROM order_source_daily_metrics
GROUP BY source_path;