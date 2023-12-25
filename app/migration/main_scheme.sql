CREATE TABLE
    "postgres"."public".ITEM
    (
        I_ID serial NOT NULL,
        ARTICUL VARCHAR NOT NULL,
        NAME VARCHAR NOT NULL,
        COUNT INTEGER DEFAULT 0 NOT NULL,
        PRICE FLOAT DEFAULT 0 NOT NULL,
        CATEGORY VARCHAR,
        SUB_CATEGORY VARCHAR,
        V_ITEM VARCHAR,
        SHOP_ID INTEGER NOT NULL, 
        PRIMARY KEY (I_ID),
        CONSTRAINT item_fk1 FOREIGN KEY (SHOP_ID) REFERENCES "postgres"."public"."shop" ("shop_id"
        )
    );

CREATE TABLE
    "postgres"."public".SALE
    (
        S_ID serial NOT NULL,
        DISCOUNT FLOAT,
        PAYMENT_TYPE INTEGER NOT NULL,
        TOTAL_SUM FLOAT NOT NULL,
        PRIMARY KEY (S_ID),
        CONSTRAINT SALE_fk1 FOREIGN KEY (PAYMENT_TYPE) REFERENCES "postgres"."public"."payment_type" ("type_id"
        )
    );

CREATE TABLE
    "postgres"."public".PAYMENT_TYPE
    (
        TYPE_ID serial NOT NULL,
        NAME VARCHAR NOT NULL,
        PRIMARY KEY (TYPE_ID)
    );

CREATE TABLE
    "postgres"."public".SALE_ITEM_LOG
    (
        ISL_ID serial NOT NULL,
        S_ID INTEGER NOT NULL,
        ITEM_ID INTEGER NOT NULL,
        DISCOUNT FLOAT,
        COUNT INTEGER NOT NULL,
        PRIMARY KEY (ISL_ID),
        CONSTRAINT ITEMSALELOG_fk1 FOREIGN KEY (S_ID) REFERENCES "postgres"."public"."sale" ("s_id"
        ),
        CONSTRAINT ITEMSALELOG_fk2 FOREIGN KEY (ITEM_ID) REFERENCES "postgres"."public"."item"
        ("i_id")
    );

CREATE TABLE
    "postgres"."public".SALE_PAYMENT_LOG
    (
        SPL_ID serial NOT NULL,
        S_ID INTEGER NOT NULL,
        PT_ID INTEGER NOT NULL,
        PAYMETN_AMOUNT FLOAT NOT NULL,
        PRIMARY KEY (SPL_ID),
        CONSTRAINT SALEPAYMENTLOG_fk1 FOREIGN KEY (S_ID) REFERENCES "postgres"."public"."sale"
        ("s_id"),
        CONSTRAINT SALEPAYMENTLOG_fk2 FOREIGN KEY (PT_ID) REFERENCES
        "postgres"."public"."payment_type" ("type_id")
    );

CREATE TABLE
    "postgres"."public".PROFILE
    (
        PROF_ID serial NOT NULL,
        LOGIN VARCHAR NOT NULL,
        PASSWORD VARCHAR NOT NULL,
        PRIMARY KEY (PROF_ID)
    );

CREATE TABLE
    "postgres"."public".SHOP
    (
        SHOP_ID serial NOT NULL,
        NAME VARCHAR NOT NULL,
        PRIMARY KEY (SHOP_ID)
    );

CREATE TABLE
    "postgres"."public".ACCESS_RIGHTS
    (
        AR_ID serial NOT NULL,
        NAME VARCHAR NOT NULL,
        PRIMARY KEY (AR_ID)
    );

CREATE TABLE
    "postgres"."public".SHOP_PROFILE
    (
        SHOP_ID INTEGER NOT NULL,
        PROF_ID INTEGER NOT NULL,
        PRIMARY KEY (SHOP_ID, PROF_ID),
        CONSTRAINT SHOPPROFILE_fk1 FOREIGN KEY (SHOP_ID) REFERENCES "postgres"."public"."shop"
        ("shop_id"),
        CONSTRAINT SHOPPROFILE_fk2 FOREIGN KEY (PROF_ID) REFERENCES "postgres"."public"."profile"
    );


CREATE TABLE
    "postgres"."public".PROFILE_AR
    (
        PROF_ID INTEGER NOT NULL,
        AR_ID INTEGER NOT NULL,
        PRIMARY KEY (PROF_ID, AR_ID),
        CONSTRAINT PROFILEAR_fk1 FOREIGN KEY (AR_ID) REFERENCES "postgres"."public"."access_rights"
        ("ar_id"),
        CONSTRAINT PROFILEAR_fk2 FOREIGN KEY (PROF_ID) REFERENCES "postgres"."public"."profile"
        ("prof_id")
    );