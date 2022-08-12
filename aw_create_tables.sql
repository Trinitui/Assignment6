DROP DATABASE IF EXISTS msds;
CREATE DATABASE msds;

DROP TABLE IF EXISTS MSDSCourseCatalog;

\c msds;

CREATE TABLE MSDSCourseCatalog (
    CID VARCHAR(100),
    CNAME VARCHAR(100) PRIMARY KEY,
    CPREREQ VARCHAR(100)

);

\c msds;

SELECT * FROM MSDSCourseCatalog;