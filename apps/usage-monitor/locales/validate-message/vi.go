package validate_message

const defaultViErrMsg = " trường không vượt qua xác thực"

var viMessages = map[string]string{
	"_": "{field}" + defaultViErrMsg, // thông điệp mặc định
	// các thông điệp tích hợp sẵn
	"_validate": "{field} không vượt qua xác nhận", // thông điệp xác nhận mặc định
	"_filter":   "Dữ liệu {field} không hợp lệ",    // lỗi lọc dữ liệu
	// giá trị int
	"min": "Giá trị tối thiểu của {field} là %v",
	"max": "Giá trị tối đa của {field} là %v",
	// kiểm tra kiểu dữ liệu: int
	"isInt":  "{field} phải là số nguyên",
	"isInt1": "{field} phải là số nguyên và giá trị tối thiểu là %d", // kiểm tra giá trị tối thiểu
	"isInt2": "{field} phải là số nguyên và trong khoảng %d - %d",    // kiểm tra giá trị tối thiểu và tối đa
	"isInts": "{field} phải là slice số nguyên",
	"isUint": "{field} phải là số nguyên không dấu (>= 0)",
	// kiểm tra kiểu dữ liệu: string
	"isString":  "{field} phải là chuỗi",
	"isString1": "{field} phải là chuỗi và độ dài tối thiểu là %d", // kiểm tra độ dài tối thiểu
	// độ dài
	"minLength": "Độ dài tối thiểu của {field} là %d",
	"maxLength": "Độ dài tối đa của {field} là %d",
	// độ dài chuỗi. tính theo rune
	"stringLength":  "Độ dài của {field} phải nằm trong khoảng %d - %d",
	"stringLength1": "Độ dài tối thiểu của {field} là %d",
	"stringLength2": "Độ dài của {field} phải nằm trong khoảng %d - %d",

	"isURL":     "{field} phải là địa chỉ URL hợp lệ",
	"isFullURL": "{field} phải là địa chỉ URL đầy đủ hợp lệ",
	"regexp":    "{field} phải phù hợp với mẫu %s",

	"isFile":  "{field} phải là tập tin đã được tải lên",
	"isImage": "{field} phải là tập tin hình ảnh đã được tải lên",

	"enum":  "{field} phải nằm trong danh sách enum %v",
	"range": "{field} phải nằm trong khoảng %d - %d",
	// So sánh giá trị
	"lt": "{field} giá trị phải nhỏ hơn %v",
	"gt": "{field} giá trị phải lớn hơn %v",
	// Yêu cầu
	"required":           "{field} là bắt buộc và không được để trống",
	"requiredIf":         "{field} là bắt buộc khi {args0} là {args1end}",
	"requiredUnless":     "{field} là bắt buộc trừ khi {args0} ở trong {args1end}",
	"requiredWith":       "{field} là bắt buộc khi {values} có mặt",
	"requiredWithAll":    "{field} là bắt buộc khi {values} có mặt",
	"requiredWithout":    "{field} là bắt buộc khi {values} không có mặt",
	"requiredWithoutAll": "{field} là bắt buộc khi không có {values} nào có mặt",
	// So sánh trường
	"eqField":  "{field} giá trị phải bằng trường %s",
	"neField":  "{field} giá trị không được bằng trường %s",
	"ltField":  "{field} giá trị phải nhỏ hơn trường %s",
	"lteField": "{field} giá trị phải nhỏ hơn hoặc bằng trường %s",
	"gtField":  "{field} giá trị phải lớn hơn trường %s",
	"gteField": "{field} giá trị phải lớn hơn hoặc bằng trường %s",
	// Loại dữ liệu
	"bool":    "{field} giá trị phải là bool",
	"float":   "{field} giá trị phải là float",
	"slice":   "{field} giá trị phải là slice",
	"map":     "{field} giá trị phải là map",
	"array":   "{field} giá trị phải là array",
	"strings": "{field} giá trị phải là []string",
	"notIn":   "{field} giá trị không được nằm trong danh sách định nghĩa %d",
	//
	"contains":    "{field} giá trị không chứa %s",
	"notContains": "{field} giá trị chứa %s",
	"startsWith":  "{field} giá trị không bắt đầu với %s",
	"endsWith":    "{field} giá trị không kết thúc với %s",
	"email":       "{field} giá trị là địa chỉ email không hợp lệ",
	"regex":       "{field} giá trị không đáp ứng kiểm tra regex",
	"file":        "{field} giá trị phải là tệp",
	"image":       "{field} giá trị phải là hình ảnh",
	// date
	"date":    "{field} giá trị phải là chuỗi ngày",
	"gtDate":  "{field} giá trị phải sau %s",
	"ltDate":  "{field} giá trị phải trước %s",
	"gteDate": "{field} giá trị phải sau hoặc bằng %s",
	"lteDate": "{field} giá trị phải trước hoặc bằng %s",
	// check char
	"hasWhitespace":  "{field} giá trị phải chứa khoảng trắng",
	"ascii":          "{field} giá trị phải là chuỗi ASCII",
	"alpha":          "{field} giá trị chỉ chứa các ký tự chữ cái",
	"alphaNum":       "{field} giá trị chỉ chứa các ký tự chữ cái và số",
	"alphaDash":      "{field} giá trị chỉ chứa các ký tự chữ cái, số, dấu gạch ngang (-) và gạch dưới (_)",
	"multiByte":      "{field} giá trị phải là chuỗi multibyte",
	"base64":         "{field} giá trị phải là chuỗi base64",
	"dnsName":        "{field} giá trị phải là chuỗi DNS",
	"dataURI":        "{field} giá trị phải là chuỗi DataURL",
	"empty":          "{field} giá trị phải là rỗng",
	"hexColor":       "{field} giá trị phải là chuỗi màu hexa",
	"hexadecimal":    "{field} giá trị phải là chuỗi hexadecimals",
	"json":           "{field} giá trị phải là chuỗi json",
	"lat":            "{field} giá trị phải là tọa độ vĩ độ",
	"lon":            "{field} giá trị phải là tọa độ kinh độ",
	"num":            "{field} giá trị phải là chuỗi số (>=0)",
	"mac":            "{field} giá trị phải là địa chỉ MAC",
	"cnMobile":       "{field} giá trị phải là chuỗi số điện thoại di động Trung Quốc 11 chữ số",
	"printableASCII": "{field} giá trị phải là chuỗi ASCII có thể in được",
	"rgbColor":       "{field} giá trị phải là chuỗi màu RGB",
	"fullURL":        "{field} giá trị phải là chuỗi đầy đủ URL",
	"full":           "{field} giá trị phải là chuỗi URL",
	"ip":             "{field} giá trị phải là chuỗi địa chỉ IP (v4 hoặc v6)",
	"ipv4":           "{field} giá trị phải là chuỗi địa chỉ IPv4",
	"ipv6":           "{field} giá trị phải là chuỗi địa chỉ IPv6",
	"CIDR":           "{field} giá trị phải là chuỗi CIDR",
	"CIDRv4":         "{field} giá trị phải là chuỗi CIDRv4",
	"CIDRv6":         "{field} giá trị phải là chuỗi CIDRv6",
	"uuid":           "{field} giá trị phải là chuỗi UUID",
	"uuid3":          "{field} giá trị phải là chuỗi UUID3",
	"uuid4":          "{field} giá trị phải là chuỗi UUID4",
	"uuid5":          "{field} giá trị phải là chuỗi UUID5",
	"filePath":       "{field} giá trị phải là đường dẫn tệp tin tồn tại",
	"unixPath":       "{field} giá trị phải là chuỗi đường dẫn unix",
	"winPath":        "{field} giá trị phải là chuỗi đường dẫn windows",
	"isbn10":         "{field} giá trị phải là chuỗi isbn10",
	"isbn13":         "{field} giá trị phải là chuỗi isbn13",
}
