package utils

import "regexp"

var RegNumber = regexp.MustCompile("[0-9A-Za-z一二三四五六七八九零壹贰叁肆伍陆柒捌玖ⅠⅡⅢⅣⅤⅥⅦⅧⅨⅩⅪⅫ⑴⑵⑶⑷⑸⑹⑺⑻⑼⑽⑽⑾⑿⒀⒁⒂⒃⒄⒅⒆⒇❶❷❸❹❺❻❼❽❾❿¹²³⁴⁵⁶⁷⁸⁹⁰𝟎𝟏𝟐𝟑𝟒𝟓𝟔𝟕𝟖𝟗𝟘𝟙𝟚𝟛𝟜𝟝𝟞𝟟𝟠𝟡𝟢𝟣𝟤𝟥𝟦𝟧𝟨𝟩𝟪𝟫𝟬𝟭𝟮𝟯𝟰𝟱𝟲𝟳𝟴𝟵𝟶𝟷𝟸𝟹𝟺𝟻𝟼𝟽𝟾𝟿₀₁₂₃₄₅₆₇₈₉①②③④⑤⑥⑦⑧⑨⑩⑪⑫⑬⑭⑮⑯⑰⑱⑲⑳㉑㉒㉓㉔㉕㉖㉗㉘㉙㉚㉛㉜㉝㉞㉟㊱㊲㊳㊴㊵㊶㊷㊸㊹㊺㊻㊼㊽㊾㊿㈠㈡㈢㈣㈤㈥㈦㈧㈨㈩⒈⒉⒊⒋⒌⒍⒎⒏⒐⒑⒒⒓⒔⒕⒖⒗⒘⒙⒚⒛㊀㊁㊂㊃㊄㊅㊆㊇㊈㊉]")
var RegEmoji = regexp.MustCompile(`[\x{1F600}-\x{1F64F}\x{1F300}-\x{1F5FF}\x{1F680}-\x{1F6FF}\x{2600}-\x{26FF}\x{2700}-\x{27BF}\x{1F900}-\x{1F9FF}]`)
