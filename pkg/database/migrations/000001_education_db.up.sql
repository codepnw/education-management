BEGIN;

-- TIME ZONE
SET TIME ZONE 'Asia/Bangkok';
-- INSTALL uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SEQUENCE users_id_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE students_id_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE teachers_id_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE courses_id_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE enroll_id_seq START WITH 1 INCREMENT BY 1;

CREATE TYPE grade_enum AS ENUM ('A', 'B', 'C', 'D', 'F');
CREATE TYPE user_role_enum AS ENUM ('admin', 'teacher', 'student');

CREATE TABLE "users" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('U', LPAD(NEXTVAL('users_id_seq')::TEXT, 6, '0')),
  "email" VARCHAR(30) NOT NULL UNIQUE,
  "username" VARCHAR(30) NOT NULL UNIQUE,
  "password" VARCHAR(30) NOT NULL,
  "role" user_role_enum NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE "oauth" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" VARCHAR NOT NULL,
  "access_token" VARCHAR NOT NULL,
  "refresh_token" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE "students" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('S', LPAD(NEXTVAL('students_id_seq')::TEXT, 6, '0')),
  "first_name" VARCHAR(50) NOT NULL,
  "last_name" VARCHAR(50) NOT NULL,
  "dob" DATE NOT NULL,
  "phone" VARCHAR NOT NULL,
  "address" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE "teachers" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('T', LPAD(NEXTVAL('teachers_id_seq')::TEXT, 6, '0')),
  "first_name" VARCHAR(50) NOT NULL,
  "last_name" VARCHAR(50) NOT NULL,
  "department" VARCHAR NOT NULL,
  "dob" DATE NOT NULL,
  "phone" VARCHAR NOT NULL,
  "address" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE "classroom" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL UNIQUE,
  "location" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE "courses" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('C', LPAD(NEXTVAL('courses_id_seq')::TEXT, 6, '0')),
  "title" VARCHAR NOT NULL,
  "description" VARCHAR NOT NULL,
  "teacher_id" VARCHAR(7) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE "enrollment" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('E', LPAD(NEXTVAL('enroll_id_seq')::TEXT, 6, '0')),
  "student_id" VARCHAR(7) NOT NULL,
  "course_id" VARCHAR(7) NOT NULL,
  "date_enrolled" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE "schedule" (
  "id" SERIAL PRIMARY KEY,
  "course_id" VARCHAR(7) NOT NULL,
  "classroom_id" INT,
  "start_time" TIME NOT NULL,
  "end_time" TIME NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE "grade" (
  "id" SERIAL PRIMARY KEY,
  "enrolledment_id" VARCHAR NOT NULL,
  "grade" grade_enum NOT NULL
);

ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE "courses" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id") ON DELETE CASCADE;
ALTER TABLE "enrollment" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id") ON DELETE CASCADE;
ALTER TABLE "enrollment" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id") ON DELETE CASCADE;
ALTER TABLE "schedule" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id") ON DELETE CASCADE;
ALTER TABLE "schedule" ADD FOREIGN KEY ("classroom_id") REFERENCES "classroom" ("id") ON DELETE CASCADE;
ALTER TABLE "grade" ADD FOREIGN KEY ("enrolledment_id") REFERENCES "enrollment" ("id") ON DELETE CASCADE;

COMMIT;