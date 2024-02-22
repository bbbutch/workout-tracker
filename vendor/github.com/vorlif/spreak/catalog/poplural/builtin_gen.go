// This file is generated by poplural/generator/main.go; DO NOT EDIT
package poplural

func init() {
	registerLanguageForm([]string{"bm", "bo", "dz", "hnj", "id", "ig", "ii", "ja", "jbo", "jv", "jw", "kde", "kea", "km", "ko", "lkt", "lo", "ms", "my", "nqo", "osa", "sah", "ses", "sg", "su", "th", "to", "tpi", "vi", "wo", "yo", "yue", "zh", "zh_Hans", "zh_Hant"},
		"nplurals=1; plural=0;", &Form{
			NPlurals: 1,
			// nplurals=1; plural=0;
			FormFunc: func(n int64) int {
				return 0
			},
		})

	registerLanguageForm([]string{"af", "an", "asa", "ast", "az", "bal", "bem", "bez", "bg", "brx", "ce", "cgg", "chr", "ckb", "da", "de", "de_AT", "de_CH", "dv", "ee", "el", "en", "en_AU", "en_CA", "en_GB", "en_US", "eo", "et", "eu", "fi", "fo", "fur", "fy", "gl", "gsw", "ha", "haw", "hu", "ia", "io", "jgo", "jmc", "ka", "kaj", "kcg", "kk", "kkj", "kl", "ks", "ksb", "ku", "ky", "lb", "lg", "lij", "mas", "mgo", "ml", "mn", "mr", "nah", "nb", "nd", "ne", "nl", "nl_BE", "nn", "nnh", "no", "nr", "ny", "nyn", "om", "or", "os", "pap", "ps", "rm", "rof", "rwk", "saq", "sc", "scn", "sd", "sdh", "seh", "sn", "so", "sq", "ss", "ssy", "st", "sv", "sw", "sw_CD", "syr", "ta", "te", "teo", "tig", "tk", "tn", "tr", "ts", "ug", "ur", "uz", "ve", "vo", "vun", "wae", "xh", "xog", "yi"},
		"nplurals=2; plural=n != 1;", &Form{
			NPlurals: 2,
			// nplurals=2; plural=n != 1;
			FormFunc: func(n int64) int {
				if n != 1 {
					return 1
				}
				return 0
			},
		})

	registerLanguageForm([]string{"ak", "am", "as", "bho", "bn", "doi", "fa", "fa_AF", "ff", "gu", "guw", "hi", "hi_Latn", "hy", "kab", "kn", "ln", "mg", "nso", "pa", "pcm", "si", "ti", "wa", "zu"},
		"nplurals=2; plural=n > 1;", &Form{
			NPlurals: 2,
			// nplurals=2; plural=n > 1;
			FormFunc: func(n int64) int {
				if n > 1 {
					return 1
				}
				return 0
			},
		})

	registerLanguageForm([]string{"ceb", "fil", "tl"},
		"nplurals=2; plural=n != 1 && n != 2 && n != 3 && ((n % 10 == 4 || n % 10 == 6) || n % 10 == 9);", &Form{
			NPlurals: 2,
			// nplurals=2; plural=n != 1 && n != 2 && n != 3 && (n % 10 == 4 || n % 10 == 6 || n % 10 == 9);
			FormFunc: func(n int64) int {
				if n != 1 && n != 2 && n != 3 && ((n%10 == 4 || n%10 == 6) || n%10 == 9) {
					return 1
				}
				return 0
			},
		})

	registerLanguageForm([]string{"is", "mk"},
		"nplurals=2; plural=(n % 10 != 1 || n % 100 == 11);", &Form{
			NPlurals: 2,
			// nplurals=2; plural=n % 10 != 1 || n % 100 == 11;
			FormFunc: func(n int64) int {
				if n%10 != 1 || n%100 == 11 {
					return 1
				}
				return 0
			},
		})

	registerLanguageForm([]string{"tzm"},
		"nplurals=2; plural=n >= 2 && (n < 11 || n > 99);", &Form{
			NPlurals: 2,
			// nplurals=2; plural=n >= 2 && (n < 11 || n > 99);
			FormFunc: func(n int64) int {
				if n >= 2 && (n < 11 || n > 99) {
					return 1
				}
				return 0
			},
		})

	registerRawForm("nplurals=2; plural=(((n == 1 || n == 2) || n == 3) || n % 10 != 4 && n % 10 != 6 && n % 10 != 9);", &Form{
		NPlurals: 2,
		// nplurals=2; plural=((n == 1 || (n == 2 || n == 3)) || n % 10 != 4 && n % 10 != 6 && n % 10 != 9);
		FormFunc: func(n int64) int {
			if ((n == 1 || n == 2) || n == 3) || n%10 != 4 && n%10 != 6 && n%10 != 9 {
				return 1
			}
			return 0
		},
	})

	registerRawForm("nplurals=2; plural=(n <= 1 || n >= 11 && n <= 99);", &Form{
		NPlurals: 2,
		// nplurals=2; plural=(n <= 1 || n >= 11 && n <= 99);
		FormFunc: func(n int64) int {
			if n <= 1 || n >= 11 && n <= 99 {
				return 1
			}
			return 0
		},
	})

	registerRawForm("nplurals=2; plural=(n == 0 || n == 1);", &Form{
		NPlurals: 2,
		// nplurals=2; plural=(n == 0 || n == 1);
		FormFunc: func(n int64) int {
			if n == 0 || n == 1 {
				return 1
			}
			return 0
		},
	})

	registerRawForm("nplurals=2; plural=n % 10 == 1 && n % 100 != 11;", &Form{
		NPlurals: 2,
		// nplurals=2; plural=n % 10 == 1 && n % 100 != 11;
		FormFunc: func(n int64) int {
			if n%10 == 1 && n%100 != 11 {
				return 1
			}
			return 0
		},
	})

	registerLanguageForm([]string{"be", "bs", "hr", "ru", "sh", "sr", "sr_ME", "uk"},
		"nplurals=3; plural=(n % 10 == 1 && n % 100 != 11) ? 0 : ((n % 10 >= 2 && n % 10 <= 4 && (n % 100 < 12 || n % 100 > 14)) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n % 10 == 1 && n % 100 != 11) ? 0 : ((n % 10 >= 2 && n % 10 <= 4 && (n % 100 < 12 || n % 100 > 14)) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n%10 == 1 && n%100 != 11 {
					return 0
				}
				if n%10 >= 2 && n%10 <= 4 && (n%100 < 12 || n%100 > 14) {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"ca", "es", "es_419", "es_ES", "es_MX", "it", "pt_PT", "vec"},
		"nplurals=3; plural=(n == 1) ? 0 : ((n != 0 && n % 1000000 == 0) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n == 1) ? 0 : ((n != 0 && n % 1000000 == 0) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n == 1 {
					return 0
				}
				if n != 0 && n%1000000 == 0 {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"cs", "sk"},
		"nplurals=3; plural=(n == 1) ? 0 : ((n >= 2 && n <= 4) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n == 1) ? 0 : ((n >= 2 && n <= 4) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n == 1 {
					return 0
				}
				if n >= 2 && n <= 4 {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"fr", "fr_CA", "fr_CH", "pt", "pt_BR"},
		"nplurals=3; plural=(n == 0 || n == 1) ? 0 : ((n != 0 && n % 1000000 == 0) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n == 0 || n == 1) ? 0 : ((n != 0 && n % 1000000 == 0) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n == 0 || n == 1 {
					return 0
				}
				if n != 0 && n%1000000 == 0 {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"he", "iu", "naq", "sat", "se", "sma", "smi", "smj", "smn", "sms"},
		"nplurals=3; plural=(n == 1) ? 0 : ((n == 2) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n == 1) ? 0 : ((n == 2) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n == 1 {
					return 0
				}
				if n == 2 {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"ksh", "lag"},
		"nplurals=3; plural=(n == 0) ? 0 : ((n == 1) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n == 0) ? 0 : ((n == 1) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n == 0 {
					return 0
				}
				if n == 1 {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"lt"},
		"nplurals=3; plural=(n % 10 == 1 && (n % 100 < 11 || n % 100 > 19)) ? 0 : ((n % 10 >= 2 && n % 10 <= 9 && (n % 100 < 11 || n % 100 > 19)) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n % 10 == 1 && (n % 100 < 11 || n % 100 > 19)) ? 0 : ((n % 10 >= 2 && n % 10 <= 9 && (n % 100 < 11 || n % 100 > 19)) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n%10 == 1 && (n%100 < 11 || n%100 > 19) {
					return 0
				}
				if n%10 >= 2 && n%10 <= 9 && (n%100 < 11 || n%100 > 19) {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"lv", "prg"},
		"nplurals=3; plural=(n % 10 == 0 || n % 100 >= 11 && n % 100 <= 19) ? 0 : ((n % 10 == 1 && n % 100 != 11) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n % 10 == 0 || n % 100 >= 11 && n % 100 <= 19) ? 0 : ((n % 10 == 1 && n % 100 != 11) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n%10 == 0 || n%100 >= 11 && n%100 <= 19 {
					return 0
				}
				if n%10 == 1 && n%100 != 11 {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"mo", "ro", "ro_MD"},
		"nplurals=3; plural=(n == 1) ? 0 : ((n == 0 || n != 1 && n % 100 >= 1 && n % 100 <= 19) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n == 1) ? 0 : ((n == 0 || n != 1 && n % 100 >= 1 && n % 100 <= 19) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n == 1 {
					return 0
				}
				if n == 0 || n != 1 && n%100 >= 1 && n%100 <= 19 {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"pl"},
		"nplurals=3; plural=(n == 1) ? 0 : ((n % 10 >= 2 && n % 10 <= 4 && (n % 100 < 12 || n % 100 > 14)) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n == 1) ? 0 : ((n % 10 >= 2 && n % 10 <= 4 && (n % 100 < 12 || n % 100 > 14)) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n == 1 {
					return 0
				}
				if n%10 >= 2 && n%10 <= 4 && (n%100 < 12 || n%100 > 14) {
					return 1
				}
				return 2
			},
		})

	registerLanguageForm([]string{"shi"},
		"nplurals=3; plural=(n == 0 || n == 1) ? 0 : ((n >= 2 && n <= 10) ? 1 : 2);", &Form{
			NPlurals: 3,
			// nplurals=3; plural=(n == 0 || n == 1) ? 0 : ((n >= 2 && n <= 10) ? 1 : 2);
			FormFunc: func(n int64) int {
				if n == 0 || n == 1 {
					return 0
				}
				if n >= 2 && n <= 10 {
					return 1
				}
				return 2
			},
		})

	registerRawForm("nplurals=3; plural=(n == 0) ? 0 : (((n == 0 || n == 1) && n != 0) ? 1 : 2);", &Form{
		NPlurals: 3,
		// nplurals=3; plural=n == 0 ? 0 : (n == 0 || n == 1) && n != 0 ? 1 : 2;
		FormFunc: func(n int64) int {
			if n == 0 {
				return 0
			}
			if (n == 0 || n == 1) && n != 0 {
				return 1
			}
			return 2
		},
	})

	registerRawForm("nplurals=3; plural=(n == 1) ? 0 : ((n == 0 || n != 1 && n % 100 >= 1 && n % 100 <= 19) ? 1 : 2);", &Form{
		NPlurals: 3,
		// nplurals=3; plural=n == 1 ? 0 : (n == 0 || n != 1 && n % 100 >= 1 && n % 100 <= 19) ? 1 : 2;
		FormFunc: func(n int64) int {
			if n == 1 {
				return 0
			}
			if n == 0 || n != 1 && n%100 >= 1 && n%100 <= 19 {
				return 1
			}
			return 2
		},
	})

	registerLanguageForm([]string{"dsb", "hsb", "sl"},
		"nplurals=4; plural=(n % 100 == 1) ? 0 : ((n % 100 == 2) ? 1 : ((n % 100 == 3 || n % 100 == 4) ? 2 : 3));", &Form{
			NPlurals: 4,
			// nplurals=4; plural=(n % 100 == 1) ? 0 : ((n % 100 == 2) ? 1 : ((n % 100 == 3 || n % 100 == 4) ? 2 : 3));
			FormFunc: func(n int64) int {
				if n%100 == 1 {
					return 0
				}
				if n%100 == 2 {
					return 1
				}
				if n%100 == 3 || n%100 == 4 {
					return 2
				}
				return 3
			},
		})

	registerLanguageForm([]string{"gd"},
		"nplurals=4; plural=(n == 1 || n == 11) ? 0 : ((n == 2 || n == 12) ? 1 : ((n >= 3 && n <= 10 || n >= 13 && n <= 19) ? 2 : 3));", &Form{
			NPlurals: 4,
			// nplurals=4; plural=(n == 1 || n == 11) ? 0 : ((n == 2 || n == 12) ? 1 : ((n >= 3 && n <= 10 || n >= 13 && n <= 19) ? 2 : 3));
			FormFunc: func(n int64) int {
				if n == 1 || n == 11 {
					return 0
				}
				if n == 2 || n == 12 {
					return 1
				}
				if n >= 3 && n <= 10 || n >= 13 && n <= 19 {
					return 2
				}
				return 3
			},
		})

	registerLanguageForm([]string{"gv"},
		"nplurals=4; plural=(n % 10 == 1) ? 0 : ((n % 10 == 2) ? 1 : (((n % 100 == 0 || (n % 100 == 20 || (n % 100 == 40 || n % 100 == 60))) || n % 100 == 80) ? 2 : 3));", &Form{
			NPlurals: 4,
			// nplurals=4; plural=(n % 10 == 1) ? 0 : ((n % 10 == 2) ? 1 : ((n % 100 == 0 || n % 100 == 20 || n % 100 == 40 || n % 100 == 60 || n % 100 == 80) ? 2 : 3));
			FormFunc: func(n int64) int {
				if n%10 == 1 {
					return 0
				}
				if n%10 == 2 {
					return 1
				}
				if (n%100 == 0 || (n%100 == 20 || (n%100 == 40 || n%100 == 60))) || n%100 == 80 {
					return 2
				}
				return 3
			},
		})

	registerRawForm("nplurals=4; plural=(n % 100 == 1) ? 0 : ((n % 100 == 2) ? 1 : ((n % 100 >= 3 && n % 100 <= 4) ? 2 : 3));", &Form{
		NPlurals: 4,
		// nplurals=4; plural=n % 100 == 1 ? 0 : n % 100 == 2 ? 1 : n % 100 >= 3 && n % 100 <= 4 ? 2 : 3;
		FormFunc: func(n int64) int {
			if n%100 == 1 {
				return 0
			}
			if n%100 == 2 {
				return 1
			}
			if n%100 >= 3 && n%100 <= 4 {
				return 2
			}
			return 3
		},
	})

	registerLanguageForm([]string{"br"},
		"nplurals=5; plural=(n % 10 == 1 && n % 100 != 11 && n % 100 != 71 && n % 100 != 91) ? 0 : ((n % 10 == 2 && n % 100 != 12 && n % 100 != 72 && n % 100 != 92) ? 1 : (((n % 10 == 3 || n % 10 == 4) || n % 10 == 9) && (n % 100 < 10 || n % 100 > 19) && (n % 100 < 70 || n % 100 > 79) && (n % 100 < 90 || n % 100 > 99) ? 2 : ((n != 0 && n % 1000000 == 0) ? 3 : 4)));", &Form{
			NPlurals: 5,
			// nplurals=5; plural=(n % 10 == 1 && n % 100 != 11 && n % 100 != 71 && n % 100 != 91) ? 0 : ((n % 10 == 2 && n % 100 != 12 && n % 100 != 72 && n % 100 != 92) ? 1 : ((((n % 10 == 3 || n % 10 == 4) || n % 10 == 9) && (n % 100 < 10 || n % 100 > 19) && (n % 100 < 70 || n % 100 > 79) && (n % 100 < 90 || n % 100 > 99)) ? 2 : ((n != 0 && n % 1000000 == 0) ? 3 : 4)));
			FormFunc: func(n int64) int {
				if n%10 == 1 && n%100 != 11 && n%100 != 71 && n%100 != 91 {
					return 0
				}
				if n%10 == 2 && n%100 != 12 && n%100 != 72 && n%100 != 92 {
					return 1
				}
				if ((n%10 == 3 || n%10 == 4) || n%10 == 9) && (n%100 < 10 || n%100 > 19) && (n%100 < 70 || n%100 > 79) && (n%100 < 90 || n%100 > 99) {
					return 2
				}
				if n != 0 && n%1000000 == 0 {
					return 3
				}
				return 4
			},
		})

	registerLanguageForm([]string{"ga"},
		"nplurals=5; plural=(n == 1) ? 0 : ((n == 2) ? 1 : ((n >= 3 && n <= 6) ? 2 : ((n >= 7 && n <= 10) ? 3 : 4)));", &Form{
			NPlurals: 5,
			// nplurals=5; plural=(n == 1) ? 0 : ((n == 2) ? 1 : ((n >= 3 && n <= 6) ? 2 : ((n >= 7 && n <= 10) ? 3 : 4)));
			FormFunc: func(n int64) int {
				if n == 1 {
					return 0
				}
				if n == 2 {
					return 1
				}
				if n >= 3 && n <= 6 {
					return 2
				}
				if n >= 7 && n <= 10 {
					return 3
				}
				return 4
			},
		})

	registerLanguageForm([]string{"mt"},
		"nplurals=5; plural=(n == 1) ? 0 : ((n == 2) ? 1 : ((n == 0 || n % 100 >= 3 && n % 100 <= 10) ? 2 : ((n % 100 >= 11 && n % 100 <= 19) ? 3 : 4)));", &Form{
			NPlurals: 5,
			// nplurals=5; plural=(n == 1) ? 0 : ((n == 2) ? 1 : ((n == 0 || n % 100 >= 3 && n % 100 <= 10) ? 2 : ((n % 100 >= 11 && n % 100 <= 19) ? 3 : 4)));
			FormFunc: func(n int64) int {
				if n == 1 {
					return 0
				}
				if n == 2 {
					return 1
				}
				if n == 0 || n%100 >= 3 && n%100 <= 10 {
					return 2
				}
				if n%100 >= 11 && n%100 <= 19 {
					return 3
				}
				return 4
			},
		})

	registerRawForm("nplurals=5; plural=(n % 10 == 1 && n % 100 != 11 && n % 100 != 71 && n % 100 != 91) ? 0 : ((n % 10 == 2 && n % 100 != 12 && n % 100 != 72 && n % 100 != 92) ? 1 : ((n % 10 >= 3 && n % 10 <= 4 || n % 10 == 9) && (n % 100 < 10 || n % 100 > 19) && (n % 100 < 70 || n % 100 > 79) && (n % 100 < 90 || n % 100 > 99) ? 2 : ((n != 0 && n % 1000000 == 0) ? 3 : 4)));", &Form{
		NPlurals: 5,
		// nplurals=5; plural=n % 10 == 1 && n % 100 != 11 && n % 100 != 71 && n % 100 != 91 ? 0 : n % 10 == 2 && n % 100 != 12 && n % 100 != 72 && n % 100 != 92 ? 1 : (n % 10 >= 3 && n % 10 <= 4 || n % 10 == 9) && (n % 100 < 10 || n % 100 > 19) && (n % 100 < 70 || n % 100 > 79) && (n % 100 < 90 || n % 100 > 99) ? 2 : n != 0 && n % 1000000 == 0 ? 3 : 4;
		FormFunc: func(n int64) int {
			if n%10 == 1 && n%100 != 11 && n%100 != 71 && n%100 != 91 {
				return 0
			}
			if n%10 == 2 && n%100 != 12 && n%100 != 72 && n%100 != 92 {
				return 1
			}
			if (n%10 >= 3 && n%10 <= 4 || n%10 == 9) && (n%100 < 10 || n%100 > 19) && (n%100 < 70 || n%100 > 79) && (n%100 < 90 || n%100 > 99) {
				return 2
			}
			if n != 0 && n%1000000 == 0 {
				return 3
			}
			return 4
		},
	})

	registerLanguageForm([]string{"ar", "ar_001", "ars"},
		"nplurals=6; plural=(n == 0) ? 0 : ((n == 1) ? 1 : ((n == 2) ? 2 : ((n % 100 >= 3 && n % 100 <= 10) ? 3 : ((n % 100 >= 11 && n % 100 <= 99) ? 4 : 5))));", &Form{
			NPlurals: 6,
			// nplurals=6; plural=(n == 0) ? 0 : ((n == 1) ? 1 : ((n == 2) ? 2 : ((n % 100 >= 3 && n % 100 <= 10) ? 3 : ((n % 100 >= 11 && n % 100 <= 99) ? 4 : 5))));
			FormFunc: func(n int64) int {
				if n == 0 {
					return 0
				}
				if n == 1 {
					return 1
				}
				if n == 2 {
					return 2
				}
				if n%100 >= 3 && n%100 <= 10 {
					return 3
				}
				if n%100 >= 11 && n%100 <= 99 {
					return 4
				}
				return 5
			},
		})

	registerLanguageForm([]string{"cy"},
		"nplurals=6; plural=(n == 0) ? 0 : ((n == 1) ? 1 : ((n == 2) ? 2 : ((n == 3) ? 3 : ((n == 6) ? 4 : 5))));", &Form{
			NPlurals: 6,
			// nplurals=6; plural=(n == 0) ? 0 : ((n == 1) ? 1 : ((n == 2) ? 2 : ((n == 3) ? 3 : ((n == 6) ? 4 : 5))));
			FormFunc: func(n int64) int {
				if n == 0 {
					return 0
				}
				if n == 1 {
					return 1
				}
				if n == 2 {
					return 2
				}
				if n == 3 {
					return 3
				}
				if n == 6 {
					return 4
				}
				return 5
			},
		})

	registerLanguageForm([]string{"kw"},
		"nplurals=6; plural=(n == 0) ? 0 : ((n == 1) ? 1 : (((((n % 100 == 2 || (n % 100 == 22 || (n % 100 == 42 || n % 100 == 62))) || n % 100 == 82) || n % 1000 == 0 && ((n % 100000 >= 1000 && n % 100000 <= 20000 || (n % 100000 == 40000 || n % 100000 == 60000)) || n % 100000 == 80000)) || n != 0 && n % 1000000 == 100000) ? 2 : (((n % 100 == 3 || (n % 100 == 23 || (n % 100 == 43 || n % 100 == 63))) || n % 100 == 83) ? 3 : ((n != 1 && ((n % 100 == 1 || (n % 100 == 21 || (n % 100 == 41 || n % 100 == 61))) || n % 100 == 81)) ? 4 : 5))));", &Form{
			NPlurals: 6,
			// nplurals=6; plural=(n == 0) ? 0 : ((n == 1) ? 1 : (((n % 100 == 2 || n % 100 == 22 || n % 100 == 42 || n % 100 == 62 || n % 100 == 82) || n % 1000 == 0 && (n % 100000 >= 1000 && n % 100000 <= 20000 || n % 100000 == 40000 || n % 100000 == 60000 || n % 100000 == 80000) || n != 0 && n % 1000000 == 100000) ? 2 : ((n % 100 == 3 || n % 100 == 23 || n % 100 == 43 || n % 100 == 63 || n % 100 == 83) ? 3 : ((n != 1 && (n % 100 == 1 || n % 100 == 21 || n % 100 == 41 || n % 100 == 61 || n % 100 == 81)) ? 4 : 5))));
			FormFunc: func(n int64) int {
				if n == 0 {
					return 0
				}
				if n == 1 {
					return 1
				}
				if (((n%100 == 2 || (n%100 == 22 || (n%100 == 42 || n%100 == 62))) || n%100 == 82) || n%1000 == 0 && ((n%100000 >= 1000 && n%100000 <= 20000 || (n%100000 == 40000 || n%100000 == 60000)) || n%100000 == 80000)) || n != 0 && n%1000000 == 100000 {
					return 2
				}
				if (n%100 == 3 || (n%100 == 23 || (n%100 == 43 || n%100 == 63))) || n%100 == 83 {
					return 3
				}
				if n != 1 && ((n%100 == 1 || (n%100 == 21 || (n%100 == 41 || n%100 == 61))) || n%100 == 81) {
					return 4
				}
				return 5
			},
		})

}
