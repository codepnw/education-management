BEGIN;

DROP TABLE IF EXISTS "students" CASCADE;
DROP TABLE IF EXISTS "teachers" CASCADE;
DROP TABLE IF EXISTS "classroom" CASCADE;
DROP TABLE IF EXISTS "courses" CASCADE;
DROP TABLE IF EXISTS "enrollment" CASCADE;
DROP TABLE IF EXISTS "schedule" CASCADE;
DROP TABLE IF EXISTS "grade" CASCADE;

DROP SEQUENCE IF EXISTS students_id_seq;
DROP SEQUENCE IF EXISTS teachers_id_seq;
DROP SEQUENCE IF EXISTS courses_id_seq;

COMMIT;