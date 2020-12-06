CREATE SCHEMA IF NOT EXISTS accounts;
SET search_path TO accounts;

CREATE TABLE IF NOT EXISTS users
  (
    USER_ID VARCHAR(120) NOT NULL,
    USER_TYPE  VARCHAR(20) NOT NULL,
    ACCOUNT_NUMBER VARCHAR(20) NOT NULL,
    INSRT_ID VARCHAR(20) NOT NULL,
    INSRT_TS TIMESTAMP NOT NULL,
    UPDT_ID VARCHAR(20) NOT NULL,
    UPDT_TS   TIMESTAMP NULL, 
    PRIMARY KEY (USER_ID)
);

CREATE TABLE "users"
(
  "user_id" text,
  "user_type" text,
  "account_number" text,
  "insrt_id" text,
  "insrt_ts" timestamp with time zone,
  "updt_id" text,
  "updt_ts" timestamp with time zone ,
  PRIMARY KEY ("user_id")
);
