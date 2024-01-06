package jobs

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"encore.app/backend/db"
	"encore.app/backend/user"
	"encore.dev/beta/auth"
	"encore.dev/rlog"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JobPreview struct {
	Id           string    `json:"id"`
	Organization string    `json:"organization"`
	Title        string    `json:"title"`
	IsPaid       bool      `json:"is_paid"`
	City         string    `json:"city"`
	Img          string    `json:"img"`
	Posted       time.Time `json:"posted"`
	Description  string    `json:"description"`
}
type JobsResponse struct {
	Jobs []JobPreview `json:"jobs"`
}

var jobsData = JobsResponse{
	Jobs: []JobPreview{
		{
			Id:           "101",
			Organization: "Global Green Tech Initiative",
			Title:        "Sustainable Software Developer",
			IsPaid:       true,
			City:         "San Francisco",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Developing and optimizing green-coded software solutions for environmental data management. Must have a passion for eco-friendly initiatives and experience in Python and cloud computing.",
		},
		{
			Id:           "102",
			Organization: "Helping Hands Tech",
			Title:        "IT Support Specialist",
			IsPaid:       true,
			City:         "Austin",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Provide tech support and troubleshooting services for our digital literacy programs. Experience with remote desktop assistance and passion for community service required.",
		},
		{
			Id:           "103",
			Organization: "TechBridge World",
			Title:        "Digital Inclusion Analyst",
			IsPaid:       true,
			City:         "New York",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Leveraging data analytics to bridge the digital divide in underserved communities. Proficiency in statistical software and a strong commitment to social equity essential.",
		},
		{
			Id:           "104",
			Organization: "HealthTech Helpers",
			Title:        "Mobile App Developer for Health",
			IsPaid:       true,
			City:         "Los Angeles",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Creating innovative mobile solutions for healthcare access. Experience with Android/iOS and a background in health informatics or public health preferred.",
		},
		{
			Id:           "105",
			Organization: "Edu-Tech Allies",
			Title:        "Educational Technology Consultant",
			IsPaid:       true,
			City:         "Chicago",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Advising schools and educational institutions on integrating technology into curricula. Requires deep knowledge of ed-tech tools and a visionary approach to learning.",
		},
		{
			Id:           "106",
			Organization: "Civic Tech Network",
			Title:        "Community Tech Coordinator",
			IsPaid:       true,
			City:         "Seattle",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Coordinating technology initiatives within community programs. Strong organizational skills and experience with volunteer management software needed.",
		},
		{
			Id:           "107",
			Organization: "Wildlife Tech Conservation",
			Title:        "Conservation Technology Specialist",
			IsPaid:       true,
			City:         "Denver",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Applying innovative tech solutions for wildlife conservation efforts. Proficiency in GIS and remote sensing technology required, along with a passion for biodiversity.",
		},
		{
			Id:           "108",
			Organization: "Non-Profit Data Analysts",
			Title:        "Social Impact Data Scientist",
			IsPaid:       true,
			City:         "Boston",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Utilizing data science techniques to measure and enhance the social impact of programs. Must be adept at Python, R, and data visualization tools.",
		},
		{
			Id:           "109",
			Organization: "Global Tech Relief",
			Title:        "Disaster Response Technologist",
			IsPaid:       true,
			City:         "Miami",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Developing and deploying tech solutions in disaster zones. Experience with rapid prototyping and knowledge of crisis management required.",
		},
		{
			Id:           "110",
			Organization: "Inclusive Tech Advocates",
			Title:        "Accessibility Tech Specialist",
			IsPaid:       true,
			City:         "Washington D.C.",
			Img:          "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPg0KPHN2ZyB3aWR0aD0iODAwcHgiIGhlaWdodD0iODAwcHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNNSA0QzQuNDQ3NzIgNCA0IDQuNDQ3NzIgNCA1VjE5QzQgMTkuNTUyMyA0LjQ0NzcyIDIwIDUgMjBIMTJIMTNDMTMuNTUyMyAyMCAxNCAxOS41NTIzIDE0IDE5VjVDMTQgNC40NDc3MiAxMy41NTIzIDQgMTMgNEg1Wk01IDIySDEySDEzSDE5QzIwLjY1NjkgMjIgMjIgMjAuNjU2OSAyMiAxOVY5QzIyIDcuMzQzMTUgMjAuNjU2OSA2IDE5IDZIMTZWNUMxNiAzLjM0MzE1IDE0LjY1NjkgMiAxMyAySDVDMy4zNDMxNSAyIDIgMy4zNDMxNSAyIDVWMTlDMiAyMC42NTY5IDMuMzQzMTUgMjIgNSAyMlpNMTkgMjBIMTUuODI5M0MxNS45Mzk4IDE5LjY4NzIgMTYgMTkuMzUwNiAxNiAxOVY4SDE5QzE5LjU1MjMgOCAyMCA4LjQ0NzcyIDIwIDlWMTlDMjAgMTkuNTUyMyAxOS41NTIzIDIwIDE5IDIwWk03IDE0SDVWMTZIN1YxNFpNOCAxNEgxMFYxNkg4VjE0Wk0xMyAxNEgxMVYxNkgxM1YxNFpNMTcgMTRIMTlWMTZIMTdWMTRaTTE5IDEwSDE3VjEySDE5VjEwWk01IDEwSDdWMTJINVYxMFpNMTAgMTBIOFYxMkgxMFYxMFpNMTEgMTBIMTNWMTJIMTFWMTBaTTcgNkg1VjhIN1Y2Wk04IDZIMTBWOEg4VjZaTTEzIDZIMTFWOEgxM1Y2WiIgZmlsbD0iIzAwMDAwMCIvPg0KPC9zdmc+",
			Posted:       time.Now(),
			Description:  "Ensuring technology products are accessible to all, including those with disabilities. Must have experience with web accessibility standards and assistive technologies.",
		},

		// Add more JobPreview structs here...
	},
}

//encore:api public method=POST path=/jobs
func Jobs(ctx context.Context) (*JobsResponse, error) {
	// Validate the email and password, for example by calling Firebase Auth: https://encore.dev/docs/how-to/firebase-auth

	rlog.Info("Fetching Jobs")
	return &jobsData, nil
}

type SaveJobRequest struct {
	JobId string `json:"jobId"`
}

//encore:api auth method=POST path=/save/:idStr
func SaveJob(ctx context.Context, idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	usr, ok := auth.Data().(*user.Data)
	if !ok {
		return fmt.Errorf("user not found")
	}
	rlog.Info("Saving Job", id, "for user", usr.Email)

	pool, err := db.Get(context.Background())
	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), "INSERT INTO saved_jobs (email, id) VALUES ($1, $2)", usr.Email, id)

	if err != nil {
		return err
	}

	return nil
}

//encore:api auth method=POST path=/unsave/:id
func UnsaveJob(ctx context.Context, id string) error {
	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	return err
	// }
	usr, ok := auth.Data().(*user.Data)
	if !ok {
		return fmt.Errorf("user not found")
	}
	rlog.Info("Unsaving Job", id, "for user", usr.Email)

	pool, err := db.Get(context.Background())
	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), "DELETE FROM saved_jobs WHERE email = $1 AND id = $2", usr.Email, id)

	if err != nil {
		return err
	}

	return nil
}

//encore:api auth method=GET path=/saved
func SavedJobs(ctx context.Context) (*JobsResponse, error) {
	usr, ok := auth.Data().(*user.Data)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	rlog.Info("Fetching Saved Jobs for user", usr.Email)

	pool, err := db.Get(context.Background())
	if err != nil {
		return nil, err
	}

	rows, err := pool.Query(context.Background(), "SELECT id FROM saved_jobs WHERE email = $1", usr.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobIds []string
	for rows.Next() {
		var jobIdStr int
		err = rows.Scan(&jobIdStr)
		if err != nil {
			return nil, err
		}
		jobId := strconv.Itoa(jobIdStr)
		jobIds = append(jobIds, jobId)
	}

	var jobs []JobPreview
	for _, jobId := range jobIds {
		for _, job := range jobsData.Jobs {
			if job.Id == jobId {
				jobs = append(jobs, job)
			}
		}
	}

	return &JobsResponse{Jobs: jobs}, nil
}
