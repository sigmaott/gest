CREATE STREAM ads_event_stream (
    id STRING,
    event STRING,
    appId STRING,
    filledDuration INTEGER
)
WITH (
    KAFKA_TOPIC='ads_event',
    VALUE_FORMAT='JSON'
);

CREATE TABLE ads_filled AS
SELECT
    appId,
    COUNT(*) AS total_ads_insert
FROM
    ads_event_stream
WHERE
        event = 'adFilled'
    WINDOW SESSION (60 SECONDS)
GROUP BY
    appId
EMIT CHANGES;