// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Language byte

const (
	LanguageEnglish             Language = 0
	LanguageFrench              Language = 1
	LanguageItalian             Language = 2
	LanguageGerman              Language = 3
	LanguageSpanish             Language = 4
	LanguageCroatian            Language = 5
	LanguageCzech               Language = 6
	LanguageDanish              Language = 7
	LanguageDutch               Language = 8
	LanguageFinnish             Language = 9
	LanguageGreek               Language = 10
	LanguageHungarian           Language = 11
	LanguageNorwegian           Language = 12
	LanguagePolish              Language = 13
	LanguagePortuguese          Language = 14
	LanguageSlovakian           Language = 15
	LanguageSlovenian           Language = 16
	LanguageSwedish             Language = 17
	LanguageRussian             Language = 18
	LanguageTurkish             Language = 19
	LanguageLatvian             Language = 20
	LanguageUkrainian           Language = 21
	LanguageArabic              Language = 22
	LanguageFarsi               Language = 23
	LanguageBulgarian           Language = 24
	LanguageRomanian            Language = 25
	LanguageChinese             Language = 26
	LanguageJapanese            Language = 27
	LanguageKorean              Language = 28
	LanguageTaiwanese           Language = 29
	LanguageThai                Language = 30
	LanguageHebrew              Language = 31
	LanguageBrazilianPortuguese Language = 32
	LanguageIndonesian          Language = 33
	LanguageMalaysian           Language = 34
	LanguageVietnamese          Language = 35
	LanguageBurmese             Language = 36
	LanguageMongolian           Language = 37
	LanguageCustom              Language = 254
	LanguageInvalid             Language = 0xFF
)

func (l Language) Byte() byte { return byte(l) }

func (l Language) String() string {
	switch l {
	case LanguageEnglish:
		return "english"
	case LanguageFrench:
		return "french"
	case LanguageItalian:
		return "italian"
	case LanguageGerman:
		return "german"
	case LanguageSpanish:
		return "spanish"
	case LanguageCroatian:
		return "croatian"
	case LanguageCzech:
		return "czech"
	case LanguageDanish:
		return "danish"
	case LanguageDutch:
		return "dutch"
	case LanguageFinnish:
		return "finnish"
	case LanguageGreek:
		return "greek"
	case LanguageHungarian:
		return "hungarian"
	case LanguageNorwegian:
		return "norwegian"
	case LanguagePolish:
		return "polish"
	case LanguagePortuguese:
		return "portuguese"
	case LanguageSlovakian:
		return "slovakian"
	case LanguageSlovenian:
		return "slovenian"
	case LanguageSwedish:
		return "swedish"
	case LanguageRussian:
		return "russian"
	case LanguageTurkish:
		return "turkish"
	case LanguageLatvian:
		return "latvian"
	case LanguageUkrainian:
		return "ukrainian"
	case LanguageArabic:
		return "arabic"
	case LanguageFarsi:
		return "farsi"
	case LanguageBulgarian:
		return "bulgarian"
	case LanguageRomanian:
		return "romanian"
	case LanguageChinese:
		return "chinese"
	case LanguageJapanese:
		return "japanese"
	case LanguageKorean:
		return "korean"
	case LanguageTaiwanese:
		return "taiwanese"
	case LanguageThai:
		return "thai"
	case LanguageHebrew:
		return "hebrew"
	case LanguageBrazilianPortuguese:
		return "brazilian_portuguese"
	case LanguageIndonesian:
		return "indonesian"
	case LanguageMalaysian:
		return "malaysian"
	case LanguageVietnamese:
		return "vietnamese"
	case LanguageBurmese:
		return "burmese"
	case LanguageMongolian:
		return "mongolian"
	case LanguageCustom:
		return "custom"
	default:
		return "LanguageInvalid(" + strconv.Itoa(int(l)) + ")"
	}
}

// FromString parse string into Language constant it's represent, return LanguageInvalid if not found.
func LanguageFromString(s string) Language {
	switch s {
	case "english":
		return LanguageEnglish
	case "french":
		return LanguageFrench
	case "italian":
		return LanguageItalian
	case "german":
		return LanguageGerman
	case "spanish":
		return LanguageSpanish
	case "croatian":
		return LanguageCroatian
	case "czech":
		return LanguageCzech
	case "danish":
		return LanguageDanish
	case "dutch":
		return LanguageDutch
	case "finnish":
		return LanguageFinnish
	case "greek":
		return LanguageGreek
	case "hungarian":
		return LanguageHungarian
	case "norwegian":
		return LanguageNorwegian
	case "polish":
		return LanguagePolish
	case "portuguese":
		return LanguagePortuguese
	case "slovakian":
		return LanguageSlovakian
	case "slovenian":
		return LanguageSlovenian
	case "swedish":
		return LanguageSwedish
	case "russian":
		return LanguageRussian
	case "turkish":
		return LanguageTurkish
	case "latvian":
		return LanguageLatvian
	case "ukrainian":
		return LanguageUkrainian
	case "arabic":
		return LanguageArabic
	case "farsi":
		return LanguageFarsi
	case "bulgarian":
		return LanguageBulgarian
	case "romanian":
		return LanguageRomanian
	case "chinese":
		return LanguageChinese
	case "japanese":
		return LanguageJapanese
	case "korean":
		return LanguageKorean
	case "taiwanese":
		return LanguageTaiwanese
	case "thai":
		return LanguageThai
	case "hebrew":
		return LanguageHebrew
	case "brazilian_portuguese":
		return LanguageBrazilianPortuguese
	case "indonesian":
		return LanguageIndonesian
	case "malaysian":
		return LanguageMalaysian
	case "vietnamese":
		return LanguageVietnamese
	case "burmese":
		return LanguageBurmese
	case "mongolian":
		return LanguageMongolian
	case "custom":
		return LanguageCustom
	default:
		return LanguageInvalid
	}
}

// List returns all constants.
func ListLanguage() []Language {
	return []Language{
		LanguageEnglish,
		LanguageFrench,
		LanguageItalian,
		LanguageGerman,
		LanguageSpanish,
		LanguageCroatian,
		LanguageCzech,
		LanguageDanish,
		LanguageDutch,
		LanguageFinnish,
		LanguageGreek,
		LanguageHungarian,
		LanguageNorwegian,
		LanguagePolish,
		LanguagePortuguese,
		LanguageSlovakian,
		LanguageSlovenian,
		LanguageSwedish,
		LanguageRussian,
		LanguageTurkish,
		LanguageLatvian,
		LanguageUkrainian,
		LanguageArabic,
		LanguageFarsi,
		LanguageBulgarian,
		LanguageRomanian,
		LanguageChinese,
		LanguageJapanese,
		LanguageKorean,
		LanguageTaiwanese,
		LanguageThai,
		LanguageHebrew,
		LanguageBrazilianPortuguese,
		LanguageIndonesian,
		LanguageMalaysian,
		LanguageVietnamese,
		LanguageBurmese,
		LanguageMongolian,
		LanguageCustom,
	}
}
