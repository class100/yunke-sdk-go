package yunke

var (
	CourseApi  *course
	LectureApi *lecture
	FileApi    *fileHandler
)

func init() {
	CourseApi = newCourse()
	LectureApi = newLecture()
	FileApi = newFileHandler()
}
