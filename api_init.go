package yunke

var (
	CourseApi  *course
	LectureApi *lecture
	FileApi    *fileHandler
	UserApi    *userHandler
)

func init() {
	CourseApi = newCourse()
	LectureApi = newLecture()
	FileApi = newFileHandler()
	UserApi = newUserHandler()
}
