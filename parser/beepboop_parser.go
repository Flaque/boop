// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 345,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 6, 2, 49, 10, 2, 13, 2, 14, 2, 50, 5, 2, 53, 10, 2, 3, 2, 3,
	2, 6, 2, 57, 10, 2, 13, 2, 14, 2, 58, 5, 2, 61, 10, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 6, 2, 67, 10, 2, 13, 2, 14, 2, 68, 5, 2, 71, 10, 2, 3, 2, 3, 2, 5,
	2, 75, 10, 2, 3, 3, 3, 3, 5, 3, 79, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 6, 4, 88, 10, 4, 13, 4, 14, 4, 89, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4, 103, 10, 4, 13, 4, 14,
	4, 104, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 111, 10, 4, 3, 5, 6, 5, 114, 10,
	5, 13, 5, 14, 5, 115, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 130, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 136,
	10, 7, 13, 7, 14, 7, 137, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5,
	7, 147, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8, 155, 10, 8, 13,
	8, 14, 8, 156, 5, 8, 159, 10, 8, 3, 8, 3, 8, 3, 8, 6, 8, 164, 10, 8, 13,
	8, 14, 8, 165, 5, 8, 168, 10, 8, 5, 8, 170, 10, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 6, 9, 178, 10, 9, 13, 9, 14, 9, 179, 5, 9, 182, 10, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 189, 10, 9, 13, 9, 14, 9, 190, 5, 9, 193,
	10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	5, 9, 206, 10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 5, 11, 219, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 6, 13, 227, 10, 13, 13, 13, 14, 13, 228, 3, 13, 5, 13, 232,
	10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 240, 10, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 6, 15, 246, 10, 15, 13, 15, 14, 15, 247, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 6, 15, 255, 10, 15, 13, 15, 14, 15, 256, 3,
	15, 3, 15, 5, 15, 261, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 267,
	10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 5, 17, 291, 10, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 7, 17, 299, 10, 17, 12, 17, 14, 17, 302, 11, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 5, 18, 308, 10, 18, 3, 18, 3, 18, 3, 18, 7, 18, 313, 10, 18,
	12, 18, 14, 18, 316, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 322, 10,
	19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 340, 10, 22, 3, 23, 3,
	23, 3, 23, 3, 23, 2, 4, 32, 34, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2, 4, 3, 2, 26, 27, 3,
	2, 30, 31, 2, 384, 2, 74, 3, 2, 2, 2, 4, 78, 3, 2, 2, 2, 6, 110, 3, 2,
	2, 2, 8, 113, 3, 2, 2, 2, 10, 129, 3, 2, 2, 2, 12, 146, 3, 2, 2, 2, 14,
	169, 3, 2, 2, 2, 16, 205, 3, 2, 2, 2, 18, 207, 3, 2, 2, 2, 20, 218, 3,
	2, 2, 2, 22, 220, 3, 2, 2, 2, 24, 231, 3, 2, 2, 2, 26, 239, 3, 2, 2, 2,
	28, 260, 3, 2, 2, 2, 30, 266, 3, 2, 2, 2, 32, 290, 3, 2, 2, 2, 34, 307,
	3, 2, 2, 2, 36, 321, 3, 2, 2, 2, 38, 323, 3, 2, 2, 2, 40, 325, 3, 2, 2,
	2, 42, 339, 3, 2, 2, 2, 44, 341, 3, 2, 2, 2, 46, 53, 7, 5, 2, 2, 47, 49,
	7, 5, 2, 2, 48, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2,
	50, 51, 3, 2, 2, 2, 51, 53, 3, 2, 2, 2, 52, 46, 3, 2, 2, 2, 52, 48, 3,
	2, 2, 2, 53, 60, 3, 2, 2, 2, 54, 61, 5, 4, 3, 2, 55, 57, 5, 4, 3, 2, 56,
	55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2,
	2, 59, 61, 3, 2, 2, 2, 60, 54, 3, 2, 2, 2, 60, 56, 3, 2, 2, 2, 61, 62,
	3, 2, 2, 2, 62, 63, 7, 2, 2, 3, 63, 75, 3, 2, 2, 2, 64, 71, 5, 4, 3, 2,
	65, 67, 5, 4, 3, 2, 66, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 66, 3,
	2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 71, 3, 2, 2, 2, 70, 64, 3, 2, 2, 2, 70,
	66, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 7, 2, 2, 3, 73, 75, 3, 2, 2,
	2, 74, 52, 3, 2, 2, 2, 74, 70, 3, 2, 2, 2, 75, 3, 3, 2, 2, 2, 76, 79, 5,
	6, 4, 2, 77, 79, 5, 10, 6, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79,
	5, 3, 2, 2, 2, 80, 81, 7, 23, 2, 2, 81, 82, 5, 44, 23, 2, 82, 83, 7, 21,
	2, 2, 83, 84, 7, 22, 2, 2, 84, 111, 3, 2, 2, 2, 85, 87, 7, 23, 2, 2, 86,
	88, 5, 44, 23, 2, 87, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 87, 3, 2,
	2, 2, 89, 90, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 7, 21, 2, 2, 92,
	93, 7, 22, 2, 2, 93, 111, 3, 2, 2, 2, 94, 95, 7, 23, 2, 2, 95, 96, 5, 44,
	23, 2, 96, 97, 7, 21, 2, 2, 97, 98, 5, 8, 5, 2, 98, 99, 7, 22, 2, 2, 99,
	111, 3, 2, 2, 2, 100, 102, 7, 23, 2, 2, 101, 103, 5, 44, 23, 2, 102, 101,
	3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2,
	2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 7, 21, 2, 2, 107, 108, 5, 8, 5, 2,
	108, 109, 7, 22, 2, 2, 109, 111, 3, 2, 2, 2, 110, 80, 3, 2, 2, 2, 110,
	85, 3, 2, 2, 2, 110, 94, 3, 2, 2, 2, 110, 100, 3, 2, 2, 2, 111, 7, 3, 2,
	2, 2, 112, 114, 5, 10, 6, 2, 113, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2,
	115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 9, 3, 2, 2, 2, 117, 130,
	5, 18, 10, 2, 118, 130, 5, 14, 8, 2, 119, 120, 5, 24, 13, 2, 120, 121,
	7, 5, 2, 2, 121, 130, 3, 2, 2, 2, 122, 123, 5, 12, 7, 2, 123, 124, 7, 5,
	2, 2, 124, 130, 3, 2, 2, 2, 125, 126, 5, 42, 22, 2, 126, 127, 7, 5, 2,
	2, 127, 130, 3, 2, 2, 2, 128, 130, 7, 5, 2, 2, 129, 117, 3, 2, 2, 2, 129,
	118, 3, 2, 2, 2, 129, 119, 3, 2, 2, 2, 129, 122, 3, 2, 2, 2, 129, 125,
	3, 2, 2, 2, 129, 128, 3, 2, 2, 2, 130, 11, 3, 2, 2, 2, 131, 132, 7, 20,
	2, 2, 132, 133, 5, 32, 17, 2, 133, 135, 7, 21, 2, 2, 134, 136, 5, 10, 6,
	2, 135, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 7, 22, 2, 2, 140, 147,
	3, 2, 2, 2, 141, 142, 7, 20, 2, 2, 142, 143, 5, 32, 17, 2, 143, 144, 7,
	21, 2, 2, 144, 145, 7, 22, 2, 2, 145, 147, 3, 2, 2, 2, 146, 131, 3, 2,
	2, 2, 146, 141, 3, 2, 2, 2, 147, 13, 3, 2, 2, 2, 148, 149, 7, 24, 2, 2,
	149, 170, 5, 26, 14, 2, 150, 151, 7, 24, 2, 2, 151, 158, 5, 26, 14, 2,
	152, 159, 7, 5, 2, 2, 153, 155, 7, 5, 2, 2, 154, 153, 3, 2, 2, 2, 155,
	156, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 159,
	3, 2, 2, 2, 158, 152, 3, 2, 2, 2, 158, 154, 3, 2, 2, 2, 159, 170, 3, 2,
	2, 2, 160, 167, 7, 24, 2, 2, 161, 168, 7, 5, 2, 2, 162, 164, 7, 5, 2, 2,
	163, 162, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165,
	166, 3, 2, 2, 2, 166, 168, 3, 2, 2, 2, 167, 161, 3, 2, 2, 2, 167, 163,
	3, 2, 2, 2, 168, 170, 3, 2, 2, 2, 169, 148, 3, 2, 2, 2, 169, 150, 3, 2,
	2, 2, 169, 160, 3, 2, 2, 2, 170, 15, 3, 2, 2, 2, 171, 172, 7, 16, 2, 2,
	172, 206, 7, 17, 2, 2, 173, 174, 7, 16, 2, 2, 174, 181, 7, 5, 2, 2, 175,
	182, 5, 18, 10, 2, 176, 178, 5, 18, 10, 2, 177, 176, 3, 2, 2, 2, 178, 179,
	3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 182, 3, 2,
	2, 2, 181, 175, 3, 2, 2, 2, 181, 177, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2,
	183, 184, 7, 17, 2, 2, 184, 206, 3, 2, 2, 2, 185, 192, 7, 16, 2, 2, 186,
	193, 5, 18, 10, 2, 187, 189, 5, 18, 10, 2, 188, 187, 3, 2, 2, 2, 189, 190,
	3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 193, 3, 2,
	2, 2, 192, 186, 3, 2, 2, 2, 192, 188, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2,
	194, 195, 7, 17, 2, 2, 195, 206, 3, 2, 2, 2, 196, 197, 7, 16, 2, 2, 197,
	198, 5, 20, 11, 2, 198, 199, 7, 17, 2, 2, 199, 206, 3, 2, 2, 2, 200, 201,
	7, 16, 2, 2, 201, 202, 7, 5, 2, 2, 202, 203, 5, 20, 11, 2, 203, 204, 7,
	17, 2, 2, 204, 206, 3, 2, 2, 2, 205, 171, 3, 2, 2, 2, 205, 173, 3, 2, 2,
	2, 205, 185, 3, 2, 2, 2, 205, 196, 3, 2, 2, 2, 205, 200, 3, 2, 2, 2, 206,
	17, 3, 2, 2, 2, 207, 208, 5, 20, 11, 2, 208, 209, 7, 5, 2, 2, 209, 19,
	3, 2, 2, 2, 210, 211, 5, 44, 23, 2, 211, 212, 7, 8, 2, 2, 212, 213, 5,
	26, 14, 2, 213, 219, 3, 2, 2, 2, 214, 215, 5, 44, 23, 2, 215, 216, 7, 8,
	2, 2, 216, 217, 5, 24, 13, 2, 217, 219, 3, 2, 2, 2, 218, 210, 3, 2, 2,
	2, 218, 214, 3, 2, 2, 2, 219, 21, 3, 2, 2, 2, 220, 221, 7, 14, 2, 2, 221,
	222, 5, 24, 13, 2, 222, 223, 7, 15, 2, 2, 223, 23, 3, 2, 2, 2, 224, 226,
	5, 44, 23, 2, 225, 227, 5, 26, 14, 2, 226, 225, 3, 2, 2, 2, 227, 228, 3,
	2, 2, 2, 228, 226, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 232, 3, 2, 2,
	2, 230, 232, 5, 44, 23, 2, 231, 224, 3, 2, 2, 2, 231, 230, 3, 2, 2, 2,
	232, 25, 3, 2, 2, 2, 233, 240, 5, 44, 23, 2, 234, 240, 5, 36, 19, 2, 235,
	240, 5, 34, 18, 2, 236, 240, 5, 22, 12, 2, 237, 240, 5, 16, 9, 2, 238,
	240, 5, 28, 15, 2, 239, 233, 3, 2, 2, 2, 239, 234, 3, 2, 2, 2, 239, 235,
	3, 2, 2, 2, 239, 236, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239, 238, 3, 2,
	2, 2, 240, 27, 3, 2, 2, 2, 241, 242, 7, 18, 2, 2, 242, 261, 7, 19, 2, 2,
	243, 245, 7, 18, 2, 2, 244, 246, 5, 30, 16, 2, 245, 244, 3, 2, 2, 2, 246,
	247, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 249,
	3, 2, 2, 2, 249, 250, 7, 19, 2, 2, 250, 261, 3, 2, 2, 2, 251, 252, 7, 18,
	2, 2, 252, 254, 7, 5, 2, 2, 253, 255, 5, 30, 16, 2, 254, 253, 3, 2, 2,
	2, 255, 256, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257,
	258, 3, 2, 2, 2, 258, 259, 7, 19, 2, 2, 259, 261, 3, 2, 2, 2, 260, 241,
	3, 2, 2, 2, 260, 243, 3, 2, 2, 2, 260, 251, 3, 2, 2, 2, 261, 29, 3, 2,
	2, 2, 262, 267, 5, 26, 14, 2, 263, 264, 5, 26, 14, 2, 264, 265, 7, 5, 2,
	2, 265, 267, 3, 2, 2, 2, 266, 262, 3, 2, 2, 2, 266, 263, 3, 2, 2, 2, 267,
	31, 3, 2, 2, 2, 268, 269, 8, 17, 1, 2, 269, 270, 5, 26, 14, 2, 270, 271,
	7, 7, 2, 2, 271, 272, 5, 26, 14, 2, 272, 291, 3, 2, 2, 2, 273, 274, 5,
	26, 14, 2, 274, 275, 7, 10, 2, 2, 275, 276, 5, 26, 14, 2, 276, 291, 3,
	2, 2, 2, 277, 278, 5, 26, 14, 2, 278, 279, 7, 11, 2, 2, 279, 280, 5, 26,
	14, 2, 280, 291, 3, 2, 2, 2, 281, 282, 5, 26, 14, 2, 282, 283, 7, 12, 2,
	2, 283, 284, 5, 26, 14, 2, 284, 291, 3, 2, 2, 2, 285, 286, 5, 26, 14, 2,
	286, 287, 7, 13, 2, 2, 287, 288, 5, 26, 14, 2, 288, 291, 3, 2, 2, 2, 289,
	291, 5, 40, 21, 2, 290, 268, 3, 2, 2, 2, 290, 273, 3, 2, 2, 2, 290, 277,
	3, 2, 2, 2, 290, 281, 3, 2, 2, 2, 290, 285, 3, 2, 2, 2, 290, 289, 3, 2,
	2, 2, 291, 300, 3, 2, 2, 2, 292, 293, 12, 5, 2, 2, 293, 294, 7, 32, 2,
	2, 294, 299, 5, 32, 17, 6, 295, 296, 12, 4, 2, 2, 296, 297, 7, 33, 2, 2,
	297, 299, 5, 32, 17, 5, 298, 292, 3, 2, 2, 2, 298, 295, 3, 2, 2, 2, 299,
	302, 3, 2, 2, 2, 300, 298, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 33, 3,
	2, 2, 2, 302, 300, 3, 2, 2, 2, 303, 304, 8, 18, 1, 2, 304, 305, 7, 27,
	2, 2, 305, 308, 5, 34, 18, 4, 306, 308, 5, 38, 20, 2, 307, 303, 3, 2, 2,
	2, 307, 306, 3, 2, 2, 2, 308, 314, 3, 2, 2, 2, 309, 310, 12, 5, 2, 2, 310,
	311, 9, 2, 2, 2, 311, 313, 5, 34, 18, 6, 312, 309, 3, 2, 2, 2, 313, 316,
	3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 35, 3, 2,
	2, 2, 316, 314, 3, 2, 2, 2, 317, 322, 5, 38, 20, 2, 318, 322, 7, 36, 2,
	2, 319, 322, 5, 40, 21, 2, 320, 322, 7, 35, 2, 2, 321, 317, 3, 2, 2, 2,
	321, 318, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 321, 320, 3, 2, 2, 2, 322,
	37, 3, 2, 2, 2, 323, 324, 7, 34, 2, 2, 324, 39, 3, 2, 2, 2, 325, 326, 9,
	3, 2, 2, 326, 41, 3, 2, 2, 2, 327, 328, 5, 26, 14, 2, 328, 329, 7, 9, 2,
	2, 329, 330, 5, 24, 13, 2, 330, 340, 3, 2, 2, 2, 331, 332, 5, 26, 14, 2,
	332, 333, 7, 9, 2, 2, 333, 334, 5, 42, 22, 2, 334, 340, 3, 2, 2, 2, 335,
	336, 5, 24, 13, 2, 336, 337, 7, 9, 2, 2, 337, 338, 5, 42, 22, 2, 338, 340,
	3, 2, 2, 2, 339, 327, 3, 2, 2, 2, 339, 331, 3, 2, 2, 2, 339, 335, 3, 2,
	2, 2, 340, 43, 3, 2, 2, 2, 341, 342, 7, 3, 2, 2, 342, 343, 7, 36, 2, 2,
	343, 45, 3, 2, 2, 2, 42, 50, 52, 58, 60, 68, 70, 74, 78, 89, 104, 110,
	115, 129, 137, 146, 156, 158, 165, 167, 169, 179, 181, 190, 192, 205, 218,
	228, 231, 239, 247, 256, 260, 266, 290, 298, 300, 307, 314, 321, 339,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "", "", "", "'=='", "'='", "'|'", "'<'", "'>'", "'<='", "'>='",
	"'('", "')'", "'{'", "'}'", "'['", "']'", "'if'", "'do'", "'end'", "'func'",
	"'return'", "'for'", "'add'", "'sub'", "'div'", "'mult'", "'true'", "'false'",
	"'or'", "'and'",
}
var symbolicNames = []string{
	"", "", "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", "PIPE",
	"LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", "LSQUIG",
	"RSQUIG", "LBLOCK", "RBLOCK", "IF", "DO", "END", "FUNC", "RETURN", "FOR",
	"PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", "OR", "AND", "INT", "QUOTED",
	"LETTERS", "STRING",
}

var ruleNames = []string{
	"boop", "code", "funcdef", "funcguts", "statement", "ifstat", "returnstat",
	"structexpr", "assignstat", "assign", "paren_fncall", "fncall", "term",
	"list", "listterm", "conditional", "math", "literal", "num", "boolexpr",
	"pipe", "label",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type BeepBoopParser struct {
	*antlr.BaseParser
}

func NewBeepBoopParser(input antlr.TokenStream) *BeepBoopParser {
	this := new(BeepBoopParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "BeepBoop.g4"

	return this
}

// BeepBoopParser tokens.
const (
	BeepBoopParserEOF          = antlr.TokenEOF
	BeepBoopParserT__0         = 1
	BeepBoopParserCOMMENT      = 2
	BeepBoopParserNEWLINE      = 3
	BeepBoopParserWHITESPACE   = 4
	BeepBoopParserEQUALS       = 5
	BeepBoopParserASSIGN       = 6
	BeepBoopParserPIPE         = 7
	BeepBoopParserLTHAN        = 8
	BeepBoopParserGTHAN        = 9
	BeepBoopParserLTHAN_EQUALS = 10
	BeepBoopParserGTHAN_EQUALS = 11
	BeepBoopParserLPAREN       = 12
	BeepBoopParserRPAREN       = 13
	BeepBoopParserLSQUIG       = 14
	BeepBoopParserRSQUIG       = 15
	BeepBoopParserLBLOCK       = 16
	BeepBoopParserRBLOCK       = 17
	BeepBoopParserIF           = 18
	BeepBoopParserDO           = 19
	BeepBoopParserEND          = 20
	BeepBoopParserFUNC         = 21
	BeepBoopParserRETURN       = 22
	BeepBoopParserFOR          = 23
	BeepBoopParserPLUS         = 24
	BeepBoopParserSUB          = 25
	BeepBoopParserDIV          = 26
	BeepBoopParserMULT         = 27
	BeepBoopParserTRUE         = 28
	BeepBoopParserFALSE        = 29
	BeepBoopParserOR           = 30
	BeepBoopParserAND          = 31
	BeepBoopParserINT          = 32
	BeepBoopParserQUOTED       = 33
	BeepBoopParserLETTERS      = 34
	BeepBoopParserSTRING       = 35
)

// BeepBoopParser rules.
const (
	BeepBoopParserRULE_boop         = 0
	BeepBoopParserRULE_code         = 1
	BeepBoopParserRULE_funcdef      = 2
	BeepBoopParserRULE_funcguts     = 3
	BeepBoopParserRULE_statement    = 4
	BeepBoopParserRULE_ifstat       = 5
	BeepBoopParserRULE_returnstat   = 6
	BeepBoopParserRULE_structexpr   = 7
	BeepBoopParserRULE_assignstat   = 8
	BeepBoopParserRULE_assign       = 9
	BeepBoopParserRULE_paren_fncall = 10
	BeepBoopParserRULE_fncall       = 11
	BeepBoopParserRULE_term         = 12
	BeepBoopParserRULE_list         = 13
	BeepBoopParserRULE_listterm     = 14
	BeepBoopParserRULE_conditional  = 15
	BeepBoopParserRULE_math         = 16
	BeepBoopParserRULE_literal      = 17
	BeepBoopParserRULE_num          = 18
	BeepBoopParserRULE_boolexpr     = 19
	BeepBoopParserRULE_pipe         = 20
	BeepBoopParserRULE_label        = 21
)

// IBoopContext is an interface to support dynamic dispatch.
type IBoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoopContext differentiates from other interfaces.
	IsBoopContext()
}

type BoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoopContext() *BoopContext {
	var p = new(BoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_boop
	return p
}

func (*BoopContext) IsBoopContext() {}

func NewBoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoopContext {
	var p = new(BoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_boop

	return p
}

func (s *BoopContext) GetParser() antlr.Parser { return s.parser }

func (s *BoopContext) EOF() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEOF, 0)
}

func (s *BoopContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(BeepBoopParserNEWLINE)
}

func (s *BoopContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, i)
}

func (s *BoopContext) AllCode() []ICodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICodeContext)(nil)).Elem())
	var tst = make([]ICodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICodeContext)
		}
	}

	return tst
}

func (s *BoopContext) Code(i int) ICodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *BoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterBoop(s)
	}
}

func (s *BoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitBoop(s)
	}
}

func (s *BoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitBoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Boop() (localctx IBoopContext) {
	localctx = NewBoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BeepBoopParserRULE_boop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(44)
				p.Match(BeepBoopParserNEWLINE)
			}

		case 2:
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(45)
						p.Match(BeepBoopParserNEWLINE)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(48)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(52)
				p.Code()
			}

		case 2:
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserIF)|(1<<BeepBoopParserFUNC)|(1<<BeepBoopParserRETURN)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserTRUE)|(1<<BeepBoopParserFALSE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32)))) != 0) {
				{
					p.SetState(53)
					p.Code()
				}

				p.SetState(56)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(60)
			p.Match(BeepBoopParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(62)
				p.Code()
			}

		case 2:
			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserIF)|(1<<BeepBoopParserFUNC)|(1<<BeepBoopParserRETURN)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserTRUE)|(1<<BeepBoopParserFALSE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32)))) != 0) {
				{
					p.SetState(63)
					p.Code()
				}

				p.SetState(66)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(70)
			p.Match(BeepBoopParserEOF)
		}

	}

	return localctx
}

// ICodeContext is an interface to support dynamic dispatch.
type ICodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCodeContext differentiates from other interfaces.
	IsCodeContext()
}

type CodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeContext() *CodeContext {
	var p = new(CodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_code
	return p
}

func (*CodeContext) IsCodeContext() {}

func NewCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeContext {
	var p = new(CodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_code

	return p
}

func (s *CodeContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeContext) CopyFrom(ctx *CodeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatementCodeContext struct {
	*CodeContext
}

func NewStatementCodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementCodeContext {
	var p = new(StatementCodeContext)

	p.CodeContext = NewEmptyCodeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CodeContext))

	return p
}

func (s *StatementCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementCodeContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementCodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterStatementCode(s)
	}
}

func (s *StatementCodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitStatementCode(s)
	}
}

func (s *StatementCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitStatementCode(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncdefCodeContext struct {
	*CodeContext
}

func NewFuncdefCodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncdefCodeContext {
	var p = new(FuncdefCodeContext)

	p.CodeContext = NewEmptyCodeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CodeContext))

	return p
}

func (s *FuncdefCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdefCodeContext) Funcdef() IFuncdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncdefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncdefContext)
}

func (s *FuncdefCodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFuncdefCode(s)
	}
}

func (s *FuncdefCodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFuncdefCode(s)
	}
}

func (s *FuncdefCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFuncdefCode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Code() (localctx ICodeContext) {
	localctx = NewCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BeepBoopParserRULE_code)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserFUNC:
		localctx = NewFuncdefCodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Funcdef()
		}

	case BeepBoopParserT__0, BeepBoopParserNEWLINE, BeepBoopParserLPAREN, BeepBoopParserLSQUIG, BeepBoopParserLBLOCK, BeepBoopParserIF, BeepBoopParserRETURN, BeepBoopParserSUB, BeepBoopParserTRUE, BeepBoopParserFALSE, BeepBoopParserINT, BeepBoopParserQUOTED, BeepBoopParserLETTERS:
		localctx = NewStatementCodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Statement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncdefContext is an interface to support dynamic dispatch.
type IFuncdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncdefContext differentiates from other interfaces.
	IsFuncdefContext()
}

type FuncdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncdefContext() *FuncdefContext {
	var p = new(FuncdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_funcdef
	return p
}

func (*FuncdefContext) IsFuncdefContext() {}

func NewFuncdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncdefContext {
	var p = new(FuncdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_funcdef

	return p
}

func (s *FuncdefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncdefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserFUNC, 0)
}

func (s *FuncdefContext) AllLabel() []ILabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILabelContext)(nil)).Elem())
	var tst = make([]ILabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILabelContext)
		}
	}

	return tst
}

func (s *FuncdefContext) Label(i int) ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *FuncdefContext) DO() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserDO, 0)
}

func (s *FuncdefContext) END() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEND, 0)
}

func (s *FuncdefContext) Funcguts() IFuncgutsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncgutsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncgutsContext)
}

func (s *FuncdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFuncdef(s)
	}
}

func (s *FuncdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFuncdef(s)
	}
}

func (s *FuncdefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFuncdef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Funcdef() (localctx IFuncdefContext) {
	localctx = NewFuncdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BeepBoopParserRULE_funcdef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Match(BeepBoopParserFUNC)
		}
		{
			p.SetState(79)
			p.Label()
		}
		{
			p.SetState(80)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(81)
			p.Match(BeepBoopParserEND)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Match(BeepBoopParserFUNC)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == BeepBoopParserT__0 {
			{
				p.SetState(84)
				p.Label()
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(89)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(90)
			p.Match(BeepBoopParserEND)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Match(BeepBoopParserFUNC)
		}
		{
			p.SetState(93)
			p.Label()
		}
		{
			p.SetState(94)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(95)
			p.Funcguts()
		}
		{
			p.SetState(96)
			p.Match(BeepBoopParserEND)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.Match(BeepBoopParserFUNC)
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == BeepBoopParserT__0 {
			{
				p.SetState(99)
				p.Label()
			}

			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(104)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(105)
			p.Funcguts()
		}
		{
			p.SetState(106)
			p.Match(BeepBoopParserEND)
		}

	}

	return localctx
}

// IFuncgutsContext is an interface to support dynamic dispatch.
type IFuncgutsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncgutsContext differentiates from other interfaces.
	IsFuncgutsContext()
}

type FuncgutsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncgutsContext() *FuncgutsContext {
	var p = new(FuncgutsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_funcguts
	return p
}

func (*FuncgutsContext) IsFuncgutsContext() {}

func NewFuncgutsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncgutsContext {
	var p = new(FuncgutsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_funcguts

	return p
}

func (s *FuncgutsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncgutsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *FuncgutsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FuncgutsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncgutsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncgutsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFuncguts(s)
	}
}

func (s *FuncgutsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFuncguts(s)
	}
}

func (s *FuncgutsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFuncguts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Funcguts() (localctx IFuncgutsContext) {
	localctx = NewFuncgutsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BeepBoopParserRULE_funcguts)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserIF)|(1<<BeepBoopParserRETURN)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserTRUE)|(1<<BeepBoopParserFALSE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32)))) != 0) {
		{
			p.SetState(110)
			p.Statement()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PipeStatementContext struct {
	*StatementContext
}

func NewPipeStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PipeStatementContext {
	var p = new(PipeStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *PipeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeStatementContext) Pipe() IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *PipeStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *PipeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterPipeStatement(s)
	}
}

func (s *PipeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitPipeStatement(s)
	}
}

func (s *PipeStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitPipeStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignStatementContext struct {
	*StatementContext
}

func NewAssignStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStatementContext {
	var p = new(AssignStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatementContext) Assignstat() IAssignstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignstatContext)
}

func (s *AssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterAssignStatement(s)
	}
}

func (s *AssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitAssignStatement(s)
	}
}

func (s *AssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type NewlineStatementContext struct {
	*StatementContext
}

func NewNewlineStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewlineStatementContext {
	var p = new(NewlineStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *NewlineStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewlineStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *NewlineStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterNewlineStatement(s)
	}
}

func (s *NewlineStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitNewlineStatement(s)
	}
}

func (s *NewlineStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitNewlineStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	*StatementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) Returnstat() IReturnstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnstatContext)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatementContext struct {
	*StatementContext
}

func NewIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementContext {
	var p = new(IfStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) Ifstat() IIfstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstatContext)
}

func (s *IfStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type FncallStatementContext struct {
	*StatementContext
}

func NewFncallStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FncallStatementContext {
	var p = new(FncallStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *FncallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FncallStatementContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *FncallStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *FncallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFncallStatement(s)
	}
}

func (s *FncallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFncallStatement(s)
	}
}

func (s *FncallStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFncallStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BeepBoopParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Assignstat()
		}

	case 2:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Returnstat()
		}

	case 3:
		localctx = NewFncallStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(117)
			p.Fncall()
		}
		{
			p.SetState(118)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 4:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)
			p.Ifstat()
		}
		{
			p.SetState(121)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 5:
		localctx = NewPipeStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.Pipe()
		}
		{
			p.SetState(124)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 6:
		localctx = NewNewlineStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(126)
			p.Match(BeepBoopParserNEWLINE)
		}

	}

	return localctx
}

// IIfstatContext is an interface to support dynamic dispatch.
type IIfstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatContext differentiates from other interfaces.
	IsIfstatContext()
}

type IfstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatContext() *IfstatContext {
	var p = new(IfstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_ifstat
	return p
}

func (*IfstatContext) IsIfstatContext() {}

func NewIfstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatContext {
	var p = new(IfstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_ifstat

	return p
}

func (s *IfstatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatContext) IF() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserIF, 0)
}

func (s *IfstatContext) Conditional() IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *IfstatContext) DO() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserDO, 0)
}

func (s *IfstatContext) END() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEND, 0)
}

func (s *IfstatContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfstatContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterIfstat(s)
	}
}

func (s *IfstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitIfstat(s)
	}
}

func (s *IfstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitIfstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Ifstat() (localctx IIfstatContext) {
	localctx = NewIfstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BeepBoopParserRULE_ifstat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Match(BeepBoopParserIF)
		}
		{
			p.SetState(130)
			p.conditional(0)
		}
		{
			p.SetState(131)
			p.Match(BeepBoopParserDO)
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserIF)|(1<<BeepBoopParserRETURN)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserTRUE)|(1<<BeepBoopParserFALSE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32)))) != 0) {
			{
				p.SetState(132)
				p.Statement()
			}

			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(137)
			p.Match(BeepBoopParserEND)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Match(BeepBoopParserIF)
		}
		{
			p.SetState(140)
			p.conditional(0)
		}
		{
			p.SetState(141)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(142)
			p.Match(BeepBoopParserEND)
		}

	}

	return localctx
}

// IReturnstatContext is an interface to support dynamic dispatch.
type IReturnstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnstatContext differentiates from other interfaces.
	IsReturnstatContext()
}

type ReturnstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnstatContext() *ReturnstatContext {
	var p = new(ReturnstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_returnstat
	return p
}

func (*ReturnstatContext) IsReturnstatContext() {}

func NewReturnstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnstatContext {
	var p = new(ReturnstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_returnstat

	return p
}

func (s *ReturnstatContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnstatContext) RETURN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserRETURN, 0)
}

func (s *ReturnstatContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ReturnstatContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(BeepBoopParserNEWLINE)
}

func (s *ReturnstatContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, i)
}

func (s *ReturnstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterReturnstat(s)
	}
}

func (s *ReturnstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitReturnstat(s)
	}
}

func (s *ReturnstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitReturnstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Returnstat() (localctx IReturnstatContext) {
	localctx = NewReturnstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BeepBoopParserRULE_returnstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Match(BeepBoopParserRETURN)
		}
		{
			p.SetState(147)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(BeepBoopParserRETURN)
		}
		{
			p.SetState(149)
			p.Term()
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(150)
				p.Match(BeepBoopParserNEWLINE)
			}

		case 2:
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(151)
						p.Match(BeepBoopParserNEWLINE)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(154)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(158)
			p.Match(BeepBoopParserRETURN)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(159)
				p.Match(BeepBoopParserNEWLINE)
			}

		case 2:
			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(160)
						p.Match(BeepBoopParserNEWLINE)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(163)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
			}

		}

	}

	return localctx
}

// IStructexprContext is an interface to support dynamic dispatch.
type IStructexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructexprContext differentiates from other interfaces.
	IsStructexprContext()
}

type StructexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructexprContext() *StructexprContext {
	var p = new(StructexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_structexpr
	return p
}

func (*StructexprContext) IsStructexprContext() {}

func NewStructexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructexprContext {
	var p = new(StructexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_structexpr

	return p
}

func (s *StructexprContext) GetParser() antlr.Parser { return s.parser }

func (s *StructexprContext) LSQUIG() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserLSQUIG, 0)
}

func (s *StructexprContext) RSQUIG() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserRSQUIG, 0)
}

func (s *StructexprContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *StructexprContext) AllAssignstat() []IAssignstatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignstatContext)(nil)).Elem())
	var tst = make([]IAssignstatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignstatContext)
		}
	}

	return tst
}

func (s *StructexprContext) Assignstat(i int) IAssignstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignstatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignstatContext)
}

func (s *StructexprContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StructexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterStructexpr(s)
	}
}

func (s *StructexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitStructexpr(s)
	}
}

func (s *StructexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitStructexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Structexpr() (localctx IStructexprContext) {
	localctx = NewStructexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BeepBoopParserRULE_structexpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Match(BeepBoopParserLSQUIG)
		}
		{
			p.SetState(170)
			p.Match(BeepBoopParserRSQUIG)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(BeepBoopParserLSQUIG)
		}
		{
			p.SetState(172)
			p.Match(BeepBoopParserNEWLINE)
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(173)
				p.Assignstat()
			}

		case 2:
			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == BeepBoopParserT__0 {
				{
					p.SetState(174)
					p.Assignstat()
				}

				p.SetState(177)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(181)
			p.Match(BeepBoopParserRSQUIG)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(183)
			p.Match(BeepBoopParserLSQUIG)
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(184)
				p.Assignstat()
			}

		case 2:
			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == BeepBoopParserT__0 {
				{
					p.SetState(185)
					p.Assignstat()
				}

				p.SetState(188)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(192)
			p.Match(BeepBoopParserRSQUIG)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.Match(BeepBoopParserLSQUIG)
		}
		{
			p.SetState(195)
			p.Assign()
		}
		{
			p.SetState(196)
			p.Match(BeepBoopParserRSQUIG)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(198)
			p.Match(BeepBoopParserLSQUIG)
		}
		{
			p.SetState(199)
			p.Match(BeepBoopParserNEWLINE)
		}
		{
			p.SetState(200)
			p.Assign()
		}
		{
			p.SetState(201)
			p.Match(BeepBoopParserRSQUIG)
		}

	}

	return localctx
}

// IAssignstatContext is an interface to support dynamic dispatch.
type IAssignstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignstatContext differentiates from other interfaces.
	IsAssignstatContext()
}

type AssignstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignstatContext() *AssignstatContext {
	var p = new(AssignstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_assignstat
	return p
}

func (*AssignstatContext) IsAssignstatContext() {}

func NewAssignstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignstatContext {
	var p = new(AssignstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_assignstat

	return p
}

func (s *AssignstatContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignstatContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *AssignstatContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *AssignstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterAssignstat(s)
	}
}

func (s *AssignstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitAssignstat(s)
	}
}

func (s *AssignstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitAssignstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Assignstat() (localctx IAssignstatContext) {
	localctx = NewAssignstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BeepBoopParserRULE_assignstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Assign()
	}
	{
		p.SetState(206)
		p.Match(BeepBoopParserNEWLINE)
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) CopyFrom(ctx *AssignContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FncallAssignContext struct {
	*AssignContext
}

func NewFncallAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FncallAssignContext {
	var p = new(FncallAssignContext)

	p.AssignContext = NewEmptyAssignContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignContext))

	return p
}

func (s *FncallAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FncallAssignContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *FncallAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserASSIGN, 0)
}

func (s *FncallAssignContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *FncallAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFncallAssign(s)
	}
}

func (s *FncallAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFncallAssign(s)
	}
}

func (s *FncallAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFncallAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAssignContext struct {
	*AssignContext
}

func NewExprAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAssignContext {
	var p = new(ExprAssignContext)

	p.AssignContext = NewEmptyAssignContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignContext))

	return p
}

func (s *ExprAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAssignContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *ExprAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserASSIGN, 0)
}

func (s *ExprAssignContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExprAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterExprAssign(s)
	}
}

func (s *ExprAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitExprAssign(s)
	}
}

func (s *ExprAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitExprAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BeepBoopParserRULE_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.Label()
		}
		{
			p.SetState(209)
			p.Match(BeepBoopParserASSIGN)
		}
		{
			p.SetState(210)
			p.Term()
		}

	case 2:
		localctx = NewFncallAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Label()
		}
		{
			p.SetState(213)
			p.Match(BeepBoopParserASSIGN)
		}
		{
			p.SetState(214)
			p.Fncall()
		}

	}

	return localctx
}

// IParen_fncallContext is an interface to support dynamic dispatch.
type IParen_fncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParen_fncallContext differentiates from other interfaces.
	IsParen_fncallContext()
}

type Paren_fncallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParen_fncallContext() *Paren_fncallContext {
	var p = new(Paren_fncallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_paren_fncall
	return p
}

func (*Paren_fncallContext) IsParen_fncallContext() {}

func NewParen_fncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Paren_fncallContext {
	var p = new(Paren_fncallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_paren_fncall

	return p
}

func (s *Paren_fncallContext) GetParser() antlr.Parser { return s.parser }

func (s *Paren_fncallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserLPAREN, 0)
}

func (s *Paren_fncallContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *Paren_fncallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserRPAREN, 0)
}

func (s *Paren_fncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Paren_fncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Paren_fncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterParen_fncall(s)
	}
}

func (s *Paren_fncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitParen_fncall(s)
	}
}

func (s *Paren_fncallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitParen_fncall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Paren_fncall() (localctx IParen_fncallContext) {
	localctx = NewParen_fncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BeepBoopParserRULE_paren_fncall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(BeepBoopParserLPAREN)
	}
	{
		p.SetState(219)
		p.Fncall()
	}
	{
		p.SetState(220)
		p.Match(BeepBoopParserRPAREN)
	}

	return localctx
}

// IFncallContext is an interface to support dynamic dispatch.
type IFncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFncallContext differentiates from other interfaces.
	IsFncallContext()
}

type FncallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFncallContext() *FncallContext {
	var p = new(FncallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_fncall
	return p
}

func (*FncallContext) IsFncallContext() {}

func NewFncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FncallContext {
	var p = new(FncallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_fncall

	return p
}

func (s *FncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FncallContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *FncallContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *FncallContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFncall(s)
	}
}

func (s *FncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFncall(s)
	}
}

func (s *FncallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFncall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Fncall() (localctx IFncallContext) {
	localctx = NewFncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BeepBoopParserRULE_fncall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.Label()
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserTRUE)|(1<<BeepBoopParserFALSE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32)))) != 0) {
			{
				p.SetState(223)
				p.Term()
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Label()
		}

	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) CopyFrom(ctx *TermContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParenfncallTermContext struct {
	*TermContext
}

func NewParenfncallTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenfncallTermContext {
	var p = new(ParenfncallTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *ParenfncallTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenfncallTermContext) Paren_fncall() IParen_fncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParen_fncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParen_fncallContext)
}

func (s *ParenfncallTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterParenfncallTerm(s)
	}
}

func (s *ParenfncallTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitParenfncallTerm(s)
	}
}

func (s *ParenfncallTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitParenfncallTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralTermContext struct {
	*TermContext
}

func NewLiteralTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralTermContext {
	var p = new(LiteralTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *LiteralTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralTermContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterLiteralTerm(s)
	}
}

func (s *LiteralTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitLiteralTerm(s)
	}
}

func (s *LiteralTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitLiteralTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

type LabelTermContext struct {
	*TermContext
}

func NewLabelTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelTermContext {
	var p = new(LabelTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *LabelTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelTermContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LabelTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterLabelTerm(s)
	}
}

func (s *LabelTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitLabelTerm(s)
	}
}

func (s *LabelTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitLabelTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructTermContext struct {
	*TermContext
}

func NewStructTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructTermContext {
	var p = new(StructTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *StructTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTermContext) Structexpr() IStructexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructexprContext)
}

func (s *StructTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterStructTerm(s)
	}
}

func (s *StructTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitStructTerm(s)
	}
}

func (s *StructTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitStructTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

type MathTermContext struct {
	*TermContext
}

func NewMathTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MathTermContext {
	var p = new(MathTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *MathTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathTermContext) Math() IMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *MathTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterMathTerm(s)
	}
}

func (s *MathTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitMathTerm(s)
	}
}

func (s *MathTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitMathTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListTermContext struct {
	*TermContext
}

func NewListTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListTermContext {
	var p = new(ListTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *ListTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTermContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ListTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterListTerm(s)
	}
}

func (s *ListTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitListTerm(s)
	}
}

func (s *ListTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitListTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BeepBoopParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLabelTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.Label()
		}

	case 2:
		localctx = NewLiteralTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
			p.Literal()
		}

	case 3:
		localctx = NewMathTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.math(0)
		}

	case 4:
		localctx = NewParenfncallTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(234)
			p.Paren_fncall()
		}

	case 5:
		localctx = NewStructTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(235)
			p.Structexpr()
		}

	case 6:
		localctx = NewListTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(236)
			p.List()
		}

	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) LBLOCK() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserLBLOCK, 0)
}

func (s *ListContext) RBLOCK() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserRBLOCK, 0)
}

func (s *ListContext) AllListterm() []IListtermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IListtermContext)(nil)).Elem())
	var tst = make([]IListtermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IListtermContext)
		}
	}

	return tst
}

func (s *ListContext) Listterm(i int) IListtermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListtermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IListtermContext)
}

func (s *ListContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BeepBoopParserRULE_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.Match(BeepBoopParserLBLOCK)
		}
		{
			p.SetState(240)
			p.Match(BeepBoopParserRBLOCK)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.Match(BeepBoopParserLBLOCK)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserTRUE)|(1<<BeepBoopParserFALSE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32)))) != 0) {
			{
				p.SetState(242)
				p.Listterm()
			}

			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(247)
			p.Match(BeepBoopParserRBLOCK)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(249)
			p.Match(BeepBoopParserLBLOCK)
		}
		{
			p.SetState(250)
			p.Match(BeepBoopParserNEWLINE)
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserTRUE)|(1<<BeepBoopParserFALSE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32)))) != 0) {
			{
				p.SetState(251)
				p.Listterm()
			}

			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(256)
			p.Match(BeepBoopParserRBLOCK)
		}

	}

	return localctx
}

// IListtermContext is an interface to support dynamic dispatch.
type IListtermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListtermContext differentiates from other interfaces.
	IsListtermContext()
}

type ListtermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListtermContext() *ListtermContext {
	var p = new(ListtermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_listterm
	return p
}

func (*ListtermContext) IsListtermContext() {}

func NewListtermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListtermContext {
	var p = new(ListtermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_listterm

	return p
}

func (s *ListtermContext) GetParser() antlr.Parser { return s.parser }

func (s *ListtermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ListtermContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *ListtermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListtermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListtermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterListterm(s)
	}
}

func (s *ListtermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitListterm(s)
	}
}

func (s *ListtermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitListterm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Listterm() (localctx IListtermContext) {
	localctx = NewListtermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BeepBoopParserRULE_listterm)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Term()
		}
		{
			p.SetState(262)
			p.Match(BeepBoopParserNEWLINE)
		}

	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) CopyFrom(ctx *ConditionalContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GthanequalsCondContext struct {
	*ConditionalContext
}

func NewGthanequalsCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GthanequalsCondContext {
	var p = new(GthanequalsCondContext)

	p.ConditionalContext = NewEmptyConditionalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalContext))

	return p
}

func (s *GthanequalsCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GthanequalsCondContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *GthanequalsCondContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *GthanequalsCondContext) GTHAN_EQUALS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserGTHAN_EQUALS, 0)
}

func (s *GthanequalsCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterGthanequalsCond(s)
	}
}

func (s *GthanequalsCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitGthanequalsCond(s)
	}
}

func (s *GthanequalsCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitGthanequalsCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualsCondContext struct {
	*ConditionalContext
}

func NewEqualsCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualsCondContext {
	var p = new(EqualsCondContext)

	p.ConditionalContext = NewEmptyConditionalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalContext))

	return p
}

func (s *EqualsCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualsCondContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *EqualsCondContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *EqualsCondContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEQUALS, 0)
}

func (s *EqualsCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterEqualsCond(s)
	}
}

func (s *EqualsCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitEqualsCond(s)
	}
}

func (s *EqualsCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitEqualsCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrCondContext struct {
	*ConditionalContext
}

func NewOrCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrCondContext {
	var p = new(OrCondContext)

	p.ConditionalContext = NewEmptyConditionalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalContext))

	return p
}

func (s *OrCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrCondContext) AllConditional() []IConditionalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionalContext)(nil)).Elem())
	var tst = make([]IConditionalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionalContext)
		}
	}

	return tst
}

func (s *OrCondContext) Conditional(i int) IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *OrCondContext) OR() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserOR, 0)
}

func (s *OrCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterOrCond(s)
	}
}

func (s *OrCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitOrCond(s)
	}
}

func (s *OrCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitOrCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type GthanCondContext struct {
	*ConditionalContext
}

func NewGthanCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GthanCondContext {
	var p = new(GthanCondContext)

	p.ConditionalContext = NewEmptyConditionalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalContext))

	return p
}

func (s *GthanCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GthanCondContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *GthanCondContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *GthanCondContext) GTHAN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserGTHAN, 0)
}

func (s *GthanCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterGthanCond(s)
	}
}

func (s *GthanCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitGthanCond(s)
	}
}

func (s *GthanCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitGthanCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndCondContext struct {
	*ConditionalContext
}

func NewAndCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndCondContext {
	var p = new(AndCondContext)

	p.ConditionalContext = NewEmptyConditionalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalContext))

	return p
}

func (s *AndCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndCondContext) AllConditional() []IConditionalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionalContext)(nil)).Elem())
	var tst = make([]IConditionalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionalContext)
		}
	}

	return tst
}

func (s *AndCondContext) Conditional(i int) IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *AndCondContext) AND() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserAND, 0)
}

func (s *AndCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterAndCond(s)
	}
}

func (s *AndCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitAndCond(s)
	}
}

func (s *AndCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitAndCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type LthanCondContext struct {
	*ConditionalContext
}

func NewLthanCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LthanCondContext {
	var p = new(LthanCondContext)

	p.ConditionalContext = NewEmptyConditionalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalContext))

	return p
}

func (s *LthanCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LthanCondContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *LthanCondContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *LthanCondContext) LTHAN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserLTHAN, 0)
}

func (s *LthanCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterLthanCond(s)
	}
}

func (s *LthanCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitLthanCond(s)
	}
}

func (s *LthanCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitLthanCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type LthanequalsCondContext struct {
	*ConditionalContext
}

func NewLthanequalsCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LthanequalsCondContext {
	var p = new(LthanequalsCondContext)

	p.ConditionalContext = NewEmptyConditionalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalContext))

	return p
}

func (s *LthanequalsCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LthanequalsCondContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *LthanequalsCondContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *LthanequalsCondContext) LTHAN_EQUALS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserLTHAN_EQUALS, 0)
}

func (s *LthanequalsCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterLthanequalsCond(s)
	}
}

func (s *LthanequalsCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitLthanequalsCond(s)
	}
}

func (s *LthanequalsCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitLthanequalsCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolCondContext struct {
	*ConditionalContext
}

func NewBoolCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolCondContext {
	var p = new(BoolCondContext)

	p.ConditionalContext = NewEmptyConditionalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalContext))

	return p
}

func (s *BoolCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolCondContext) Boolexpr() IBoolexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolexprContext)
}

func (s *BoolCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterBoolCond(s)
	}
}

func (s *BoolCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitBoolCond(s)
	}
}

func (s *BoolCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitBoolCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Conditional() (localctx IConditionalContext) {
	return p.conditional(0)
}

func (p *BeepBoopParser) conditional(_p int) (localctx IConditionalContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewConditionalContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionalContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, BeepBoopParserRULE_conditional, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEqualsCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(267)
			p.Term()
		}
		{
			p.SetState(268)
			p.Match(BeepBoopParserEQUALS)
		}
		{
			p.SetState(269)
			p.Term()
		}

	case 2:
		localctx = NewLthanCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(271)
			p.Term()
		}
		{
			p.SetState(272)
			p.Match(BeepBoopParserLTHAN)
		}
		{
			p.SetState(273)
			p.Term()
		}

	case 3:
		localctx = NewGthanCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(275)
			p.Term()
		}
		{
			p.SetState(276)
			p.Match(BeepBoopParserGTHAN)
		}
		{
			p.SetState(277)
			p.Term()
		}

	case 4:
		localctx = NewLthanequalsCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(279)
			p.Term()
		}
		{
			p.SetState(280)
			p.Match(BeepBoopParserLTHAN_EQUALS)
		}
		{
			p.SetState(281)
			p.Term()
		}

	case 5:
		localctx = NewGthanequalsCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(283)
			p.Term()
		}
		{
			p.SetState(284)
			p.Match(BeepBoopParserGTHAN_EQUALS)
		}
		{
			p.SetState(285)
			p.Term()
		}

	case 6:
		localctx = NewBoolCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(287)
			p.Boolexpr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(296)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOrCondContext(p, NewConditionalContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_conditional)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(291)
					p.Match(BeepBoopParserOR)
				}
				{
					p.SetState(292)
					p.conditional(4)
				}

			case 2:
				localctx = NewAndCondContext(p, NewConditionalContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_conditional)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(294)
					p.Match(BeepBoopParserAND)
				}
				{
					p.SetState(295)
					p.conditional(3)
				}

			}

		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}

	return localctx
}

// IMathContext is an interface to support dynamic dispatch.
type IMathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathContext differentiates from other interfaces.
	IsMathContext()
}

type MathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathContext() *MathContext {
	var p = new(MathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_math
	return p
}

func (*MathContext) IsMathContext() {}

func NewMathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathContext {
	var p = new(MathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_math

	return p
}

func (s *MathContext) GetParser() antlr.Parser { return s.parser }

func (s *MathContext) CopyFrom(ctx *MathContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SoloMathContext struct {
	*MathContext
}

func NewSoloMathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SoloMathContext {
	var p = new(SoloMathContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *SoloMathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoloMathContext) Num() INumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *SoloMathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterSoloMath(s)
	}
}

func (s *SoloMathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitSoloMath(s)
	}
}

func (s *SoloMathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitSoloMath(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditiveMathContext struct {
	*MathContext
	op antlr.Token
}

func NewAdditiveMathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveMathContext {
	var p = new(AdditiveMathContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *AdditiveMathContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveMathContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveMathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveMathContext) AllMath() []IMathContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathContext)(nil)).Elem())
	var tst = make([]IMathContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathContext)
		}
	}

	return tst
}

func (s *AdditiveMathContext) Math(i int) IMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *AdditiveMathContext) PLUS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserPLUS, 0)
}

func (s *AdditiveMathContext) SUB() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserSUB, 0)
}

func (s *AdditiveMathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterAdditiveMath(s)
	}
}

func (s *AdditiveMathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitAdditiveMath(s)
	}
}

func (s *AdditiveMathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitAdditiveMath(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnarySubMathContext struct {
	*MathContext
}

func NewUnarySubMathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnarySubMathContext {
	var p = new(UnarySubMathContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *UnarySubMathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarySubMathContext) SUB() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserSUB, 0)
}

func (s *UnarySubMathContext) Math() IMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *UnarySubMathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterUnarySubMath(s)
	}
}

func (s *UnarySubMathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitUnarySubMath(s)
	}
}

func (s *UnarySubMathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitUnarySubMath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Math() (localctx IMathContext) {
	return p.math(0)
}

func (p *BeepBoopParser) math(_p int) (localctx IMathContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMathContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, BeepBoopParserRULE_math, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserSUB:
		localctx = NewUnarySubMathContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(302)
			p.Match(BeepBoopParserSUB)
		}
		{
			p.SetState(303)
			p.math(2)
		}

	case BeepBoopParserINT:
		localctx = NewSoloMathContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(304)
			p.Num()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveMathContext(p, NewMathContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_math)
			p.SetState(307)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(308)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AdditiveMathContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == BeepBoopParserPLUS || _la == BeepBoopParserSUB) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AdditiveMathContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(309)
				p.math(4)
			}

		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QuotedLiteralContext struct {
	*LiteralContext
}

func NewQuotedLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedLiteralContext {
	var p = new(QuotedLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *QuotedLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedLiteralContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserQUOTED, 0)
}

func (s *QuotedLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterQuotedLiteral(s)
	}
}

func (s *QuotedLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitQuotedLiteral(s)
	}
}

func (s *QuotedLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitQuotedLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumLiteralContext struct {
	*LiteralContext
}

func NewNumLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumLiteralContext {
	var p = new(NumLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *NumLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumLiteralContext) Num() INumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *NumLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterNumLiteral(s)
	}
}

func (s *NumLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitNumLiteral(s)
	}
}

func (s *NumLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitNumLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type LettersLiteralContext struct {
	*LiteralContext
}

func NewLettersLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LettersLiteralContext {
	var p = new(LettersLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LettersLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LettersLiteralContext) LETTERS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserLETTERS, 0)
}

func (s *LettersLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterLettersLiteral(s)
	}
}

func (s *LettersLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitLettersLiteral(s)
	}
}

func (s *LettersLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitLettersLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolLiteralContext struct {
	*LiteralContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) Boolexpr() IBoolexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolexprContext)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitBoolLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BeepBoopParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(319)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserINT:
		localctx = NewNumLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.Num()
		}

	case BeepBoopParserLETTERS:
		localctx = NewLettersLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(316)
			p.Match(BeepBoopParserLETTERS)
		}

	case BeepBoopParserTRUE, BeepBoopParserFALSE:
		localctx = NewBoolLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(317)
			p.Boolexpr()
		}

	case BeepBoopParserQUOTED:
		localctx = NewQuotedLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(318)
			p.Match(BeepBoopParserQUOTED)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_num
	return p
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) INT() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserINT, 0)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitNum(s)
	}
}

func (s *NumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Num() (localctx INumContext) {
	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BeepBoopParserRULE_num)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(BeepBoopParserINT)
	}

	return localctx
}

// IBoolexprContext is an interface to support dynamic dispatch.
type IBoolexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolexprContext differentiates from other interfaces.
	IsBoolexprContext()
}

type BoolexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolexprContext() *BoolexprContext {
	var p = new(BoolexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_boolexpr
	return p
}

func (*BoolexprContext) IsBoolexprContext() {}

func NewBoolexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolexprContext {
	var p = new(BoolexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_boolexpr

	return p
}

func (s *BoolexprContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolexprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserTRUE, 0)
}

func (s *BoolexprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserFALSE, 0)
}

func (s *BoolexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterBoolexpr(s)
	}
}

func (s *BoolexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitBoolexpr(s)
	}
}

func (s *BoolexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitBoolexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Boolexpr() (localctx IBoolexprContext) {
	localctx = NewBoolexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BeepBoopParserRULE_boolexpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BeepBoopParserTRUE || _la == BeepBoopParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPipeContext is an interface to support dynamic dispatch.
type IPipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeContext differentiates from other interfaces.
	IsPipeContext()
}

type PipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeContext() *PipeContext {
	var p = new(PipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_pipe
	return p
}

func (*PipeContext) IsPipeContext() {}

func NewPipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeContext {
	var p = new(PipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_pipe

	return p
}

func (s *PipeContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *PipeContext) PIPE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserPIPE, 0)
}

func (s *PipeContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *PipeContext) Pipe() IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *PipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterPipe(s)
	}
}

func (s *PipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitPipe(s)
	}
}

func (s *PipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitPipe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Pipe() (localctx IPipeContext) {
	localctx = NewPipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BeepBoopParserRULE_pipe)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(325)
			p.Term()
		}
		{
			p.SetState(326)
			p.Match(BeepBoopParserPIPE)
		}
		{
			p.SetState(327)
			p.Fncall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Term()
		}
		{
			p.SetState(330)
			p.Match(BeepBoopParserPIPE)
		}
		{
			p.SetState(331)
			p.Pipe()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(333)
			p.Fncall()
		}
		{
			p.SetState(334)
			p.Match(BeepBoopParserPIPE)
		}
		{
			p.SetState(335)
			p.Pipe()
		}

	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) LETTERS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserLETTERS, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BeepBoopParserRULE_label)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(BeepBoopParserT__0)
	}
	{
		p.SetState(340)
		p.Match(BeepBoopParserLETTERS)
	}

	return localctx
}

func (p *BeepBoopParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *ConditionalContext = nil
		if localctx != nil {
			t = localctx.(*ConditionalContext)
		}
		return p.Conditional_Sempred(t, predIndex)

	case 16:
		var t *MathContext = nil
		if localctx != nil {
			t = localctx.(*MathContext)
		}
		return p.Math_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *BeepBoopParser) Conditional_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *BeepBoopParser) Math_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
