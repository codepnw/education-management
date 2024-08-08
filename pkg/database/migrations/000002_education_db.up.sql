BEGIN;

INSERT INTO "users" (
	"email",
	"username",
	"password",
	"role"
) VALUES ('admin1@mail.com', 'admin', '123123', 'admin');

INSERT INTO "teachers" (
    "first_name",
    "last_name",
    "department",
    "dob",
    "phone",
    "address"
) VALUES ('pensri', 'wangdee', 'english', '1990-02-10', '0612354395', 'pattaya, thailand');

INSERT INTO "students" (
    "first_name",
    "last_name",
    "dob",
    "phone",
    "address"
) VALUES ('chookchai', 'jaidee', '2010-12-20', '0612354395', 'london, england');

INSERT INTO "courses" (
    "title",
    "description",
    "teacher_id"
) VALUES ('english 101', 'learn basic english noob to pro', 'T000001');

INSERT INTO "classroom" (
    "name",
    "location"
) VALUES ('abc001', 'Bangkok-Thailand');

INSERT INTO "enrollment" (
    "student_id",
    "course_id",
    "date_enrolled"
) VALUES ('S000001', 'C000001', '2024-06-23 16:30');

insert into "schedule" (
	"course_id",
	"classroom_id",
	"start_time",
	"end_time"
) values ('C000001', '1', '09:00', '12:00');

INSERT INTO "grade" (
    "enrolledment_id",
    "grade"
) VALUES ('E000001', 'F');

COMMIT;