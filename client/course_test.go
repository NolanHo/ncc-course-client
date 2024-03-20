package client

import "testing"

func TestClient_GetCourses(t *testing.T) {
	c := NewClient()
	c.SetToken("eyJhbGciOiJIUzUxMiJ9.eyJsb2dpbl91c2VyX2tleSI6IjhlMDFmNWMxLTJkOTAtNDBkNi1hOTIwLTY5MjMxNzg3NTNhZiJ9.gCstFDqgDcks_fEqOJlkDqcIwtPbrX7GnlRbA0l-MjjBffLLuHppllgs9ag1HhCdDsSwZxuW2MC6oNmA65IpOw")
	courses, err := c.GetCourses()
	if err != nil {
		t.Errorf("GetCourses() failed: %v", err)
		return
	}

	t.Logf("GetCourses() success: courses=%+v", courses)
}

func TestClient_SelectCourse(t *testing.T) {
	c := NewClient()
	c.SetToken("eyJhbGciOiJIUzUxMiJ9.eyJsb2dpbl91c2VyX2tleSI6IjhlMDFmNWMxLTJkOTAtNDBkNi1hOTIwLTY5MjMxNzg3NTNhZiJ9.gCstFDqgDcks_fEqOJlkDqcIwtPbrX7GnlRbA0l-MjjBffLLuHppllgs9ag1HhCdDsSwZxuW2MC6oNmA65IpOw")
	course := Course{
		CourseId: 1089,
	}
	err := c.SelectCourse(course)
	if err != nil {
		t.Errorf("SelectCourse() failed: %v", err)
		return
	}

	t.Logf("SelectCourse() success")
}
