CREATE TABLE logs
(
    date    Date,
    size    Int32,
    message String
) ENGINE = MergeTree(date, message, 8192);
