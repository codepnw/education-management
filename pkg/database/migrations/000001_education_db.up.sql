BEGIN;

-- TIME ZONE
SET TIME ZONE 'Asia/Bangkok';
-- INSTALL uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SEQUENCE students_id_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE teachers_id_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE courses_id_seq START WITH 1 INCREMENT BY 1;

CREATE TABLE "students" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('S', LPAD(NEXTVAL('students_id_seq')::TEXT, 6, '0')),
  "first_name" VARCHAR(50) NOT NULL,
  "last_name" VARCHAR(50) NOT NULL,
  "dob" DATE NOT NULL,
  "phone" VARCHAR NOT NULL,
  "address" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "teachers" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('T', LPAD(NEXTVAL('teachers_id_seq')::TEXT, 6, '0')),
  "first_name" VARCHAR(50) NOT NULL,
  "last_name" VARCHAR(50) NOT NULL,
  "department" VARCHAR NOT NULL,
  "dob" DATE NOT NULL,
  "phone" VARCHAR NOT NULL,
  "address" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "classroom" (
  "id" VARCHAR PRIMARY KEY,
  "name" VARCHAR NOT NULL UNIQUE,
  "location" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "courses" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('C', LPAD(NEXTVAL('courses_id_seq')::TEXT, 6, '0')),
  "title" VARCHAR NOT NULL,
  "description" VARCHAR NOT NULL,
  "teacher_id" VARCHAR(7) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "enrollment" (
  "id" SERIAL PRIMARY KEY,
  "student_id" VARCHAR(7) NOT NULL,
  "course_id" VARCHAR(7) NOT NULL,
  "date_enrolled" TIMESTAMP NOT NULL DEFAULT now(),
  "grade" VARCHAR,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "schedule" (
  "id" SERIAL PRIMARY KEY,
  "course_id" VARCHAR(7) NOT NULL,
  "classroom_id" VARCHAR,
  "start_time" VARCHAR NOT NULL,
  "end_time" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "grade" (
  "enrolledment_id" INT PRIMARY KEY,
  "grade" CHAR(1)
);

ALTER TABLE "courses" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id");
ALTER TABLE "enrollment" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");
ALTER TABLE "enrollment" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");
ALTER TABLE "schedule" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");
ALTER TABLE "schedule" ADD FOREIGN KEY ("classroom_id") REFERENCES "classroom" ("id");
ALTER TABLE "grade" ADD FOREIGN KEY ("enrolledment_id") REFERENCES "enrollment" ("id");

COMMIT;