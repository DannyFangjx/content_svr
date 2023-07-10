package validator

var Uint32Schema = map[string]interface{}{
	"type":    "integer",
	"minimum": 0,
	"maximum": uint32((1 << 32) - 1),
}

var Uint8Schema = map[string]interface{}{
	"type":    "integer",
	"minimum": 0,
	"maximum": 255,
}

var Uint64Schema = map[string]interface{}{
	"type":    "integer",
	"minimum": 0,
	"maximum": uint64((1 << 64) - 1),
}

var Int64Schema = map[string]interface{}{
	"type": "integer",
}

var Uint64NotEmptySchema = map[string]interface{}{
	"type":    "integer",
	"minimum": 1,
	"maximum": uint64((1 << 64) - 1),
}

var Uint64NotBlankSchema = map[string]interface{}{
	"type":    "integer",
	"minimum": 1,
	"maximum": uint64((1 << 64) - 1),
}

var StringSchema = map[string]interface{}{
	"type": "string",
}

var BusinessLicenseSchema = map[string]interface{}{
	"type":      "string",
	"minLength": 0,
	"maxLength": 40,
	"pattern":   `^$|^[0-9a-zA-Z]{0,40}$`,
}

var String0To20Schema = map[string]interface{}{
	"type":      "string",
	"minLength": 0,
	"maxLength": 20,
}

var String0To100Schema = map[string]interface{}{
	"type":      "string",
	"minLength": 0,
	"maxLength": 100,
}

var String0To700Schema = map[string]interface{}{
	"type":      "string",
	"minLength": 0,
	"maxLength": 700,
}

var BrandNameSchema = map[string]interface{}{
	"type":      "string",
	"minLength": 0,
	"maxLength": 70,
}

var String100Schema = map[string]interface{}{
	"type":      "string",
	"minLength": 1,
	"maxLength": 100,
}

var Str128Schema = map[string]interface{}{
	"type":      "string",
	"minLength": 1,
	"maxLength": 128,
}

var Str256Schema = map[string]interface{}{
	"type":      "string",
	"minLength": 1,
	"maxLength": 256,
}

var Str1000Schema = map[string]interface{}{
	"type":      "string",
	"minLength": 0,
	"maxLength": 1000,
}

var Str100Schema = map[string]interface{}{
	"type":      "string",
	"minLength": 0,
	"maxLength": 100,
}

var StrNotBlankSchema = map[string]interface{}{
	"type":      "string",
	"minLength": 1,
}

var BooleanSchema = map[string]interface{}{
	"type": "boolean",
}

var UrlSchema = map[string]interface{}{
	"type":      "string",
	"minLength": 5,
	"maxLength": 2000,
	"format":    `^(http|ftp|https):\/\/[\w\-_]+(\.[\w\-\/_!]+)+([\w\-\.\$ ,!@?^=%&amp;:/~\+#\(\)\|]*[\w\-\/!@?^=%&amp;/~\+#])?$`,
}

var FloatSchema = map[string]interface{}{
	"type": "number",
}

var UFloatSchema = map[string]interface{}{
	"type":    "number",
	"minimum": 0,
}

var DateSchema = map[string]interface{}{
	"type":    "string",
	"pattern": `^\d{4}-\d{2}-\d{2}$`,
}

var IntsStrSchema = map[string]interface{}{
	"type":    "string",
	"pattern": "^\\d+(\\,\\d+)*$",
}

var IntStrSchema = map[string]interface{}{
	"type":    "string",
	"pattern": "^\\d+$",
}

var SizeSchema = map[string]interface{}{
	"type":    "integer",
	"minimum": 0,
	"maximum": 100,
}

var PageSchema = map[string]interface{}{
	"type":    "integer",
	"minimum": 1,
}

var EmailSchema = map[string]interface{}{
	"type": "string",
	"pattern": "^[-a-zA-Z0-9~!$%^&*_=+}{\\'?]+(\\.[-a-zA-Z0-9~!$%^&*_=+}{\\'?]+)*@([a-zA-Z0-9_]" +
		"[-a-zA-Z0-9_]*(\\.[-a-zA-Z0-9_]+)*\\.(aero|asia|arpa|world|xxx|biz|com|coop|edu|email|gov|info|" +
		"int|mil|museum|name|net|org|pro|travel|mobi|[a-zA-Z]{1,3})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}" +
		"\\.[0-9]{1,3}))(:[0-9]{1,5})?$",
	"unique": true,
}

var EnableEmptyEmailSchema = map[string]interface{}{
	"type": "string",
	"pattern": "^$|^[-a-zA-Z0-9~!$%^&*_=+}{\\'?]+(\\.[-a-zA-Z0-9~!$%^&*_=+}{\\'?]+)*@([a-zA-Z0-9_]" +
		"[-a-zA-Z0-9_]*(\\.[-a-zA-Z0-9_]+)*\\.(aero|asia|arpa|world|xxx|biz|com|coop|edu|email|gov|info|" +
		"int|mil|museum|name|net|org|pro|travel|mobi|[a-zA-Z]{1,3})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}" +
		"\\.[0-9]{1,3}))(:[0-9]{1,5})?$",
}

var EnableEmptyPhoneSchema = map[string]interface{}{
	"type":    "string",
	"pattern": "^$|^[0-9]{7,11}$",
}

var EmailSchemaAllowBlank = map[string]interface{}{
	"type": "string",
	"pattern": "^[-a-zA-Z0-9~!$%^&*_=+}{\\'?]+(\\.[-a-zA-Z0-9~!$%^&*_=+}{\\'?]+)*@([a-zA-Z0-9_]" +
		"[-a-zA-Z0-9_]*(\\.[-a-zA-Z0-9_]+)*\\.(aero|asia|arpa|world|xxx|biz|com|coop|edu|email|gov|info|" +
		"int|mil|museum|name|net|org|pro|travel|mobi|[a-zA-Z]{1,3})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}" +
		"\\.[0-9]{1,3}))(:[0-9]{1,5})?$",
	"unique":      true,
	"allow_blank": true,
}

var EmailSchemaAllowBlankNotUnique = map[string]interface{}{
	"type": "string",
	"pattern": "^[-a-zA-Z0-9~!$%^&*_=+}{\\'?]+(\\.[-a-zA-Z0-9~!$%^&*_=+}{\\'?]+)*@([a-zA-Z0-9_]" +
		"[-a-zA-Z0-9_]*(\\.[-a-zA-Z0-9_]+)*\\.(aero|asia|arpa|world|xxx|biz|com|coop|edu|email|gov|info|" +
		"int|mil|museum|name|net|org|pro|travel|mobi|[a-zA-Z]{1,3})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}" +
		"\\.[0-9]{1,3}))(:[0-9]{1,5})?$",
	"allow_blank": true,
}

var CBPhoneSchema = map[string]interface{}{
	"type": "string",
	"pattern": "^((\\+886[0-9]{9,10})|(\\+62[0-9]{9,13})|(\\+60[0-9]{9,11})|(\\+84[0-9]{9,11})" +
		"|(\\+66[0-9]{9,10})|(\\+63[0-9]{10,11})|(\\+86[0-9]{11,12})|(\\+852[0-9]{8,9})" +
		"|(\\+82[0-9]{9,13})|(\\+55[0-9]{10,11})|(\\+65[0-9]{8,9})){1,1}$",
	"unique": true,
}

var MerchantPhoneSchema = map[string]interface{}{
	"type":    "string",
	"pattern": "^$|^[0-9]{7,11}$",
	"unique":  false,
}

var CBPhoneCodeSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"+886", "+62", "+60", "+84", "+66", "+63", "+86", "+852", "+65", "+82", "+55"},
}

var PhoneCodeSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"", "+86", "+886", "+852", "+853", "+65", "+62", "+60", "+84",
		"+66", "+63", "+82", "+81", "+001", "+44", "+49", "+39", "+33", "+64",
		"+61", "+90", "+880", "+31", "+41", "+34", "+32", "+56", "+57", "+55", "+52", "+48"},
}

var RegionSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"id", "my", "ph", "sg", "th", "tw", "vn", "cn", "hk", "kr", "br"},
}

var ProductCatSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"", "Accessories", "Beauty & Personal Care", "Computers & Laptops", "Design & Crafts", "Food", "Hobbies, Games & Collectibles", "Home & Living",
		"Home Appliances", "Men's Bags & Wallets", "Men's Clothing", "Men's Shoes", "Mobile & Gadgets", "Muslim Wear", "Sports & Outdoor", "Tickets & Vouchers", "Toys, Kids & Babies",
		"Watches", "Women's Bags & Purses", "Women's Clothing", "Women's Shoes", "Bags", "Books", "Pet Accessories", "Automotive Parts & Accessories", "Everything Else"},
}

var SellerAttributeSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"", "Own Factory", "None"},
}

var LeadSizeSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"", "small", "medium", "big"},
}

var CompanySizeSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"", "<10", "10-30", "31-50", "51-100", "101-200", "201-500", ">500"},
}

var SalesPotentialSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"", "1-10", "10-50", "50-100", "100-500", "500-1000", "1000-5000", "5000-10000", ">10000"},
}

var SKUPotentialSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"", "<50", "50-100", "101-200", "201-1000", ">1000"},
}

var AvgUnitPriceSchema = map[string]interface{}{
	"type": "string",
	"enum": []string{"", "1-10", "10-20", "20-50", "50-100", "100-500", ">500"},
}

var Int64ArrSchemaNotBlank = map[string]interface{}{
	"type":     "array",
	"items":    Int64Schema,
	"minItems": 1,    // >=0
	"maxItems": 5000, // <=3
}
