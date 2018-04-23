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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 392,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3, 2, 6,
	2, 57, 10, 2, 13, 2, 14, 2, 58, 5, 2, 61, 10, 2, 3, 2, 3, 2, 6, 2, 65,
	10, 2, 13, 2, 14, 2, 66, 5, 2, 69, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2,
	75, 10, 2, 13, 2, 14, 2, 76, 5, 2, 79, 10, 2, 3, 2, 3, 2, 5, 2, 83, 10,
	2, 3, 3, 3, 3, 5, 3, 87, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 6, 4, 96, 10, 4, 13, 4, 14, 4, 97, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4, 111, 10, 4, 13, 4, 14, 4, 112, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 119, 10, 4, 3, 5, 6, 5, 122, 10, 5, 13, 5, 14,
	5, 123, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 142, 10, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 5, 7, 151, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 156, 10,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 162, 10, 9, 13, 9, 14, 9, 163, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 173, 10, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 6, 10, 181, 10, 10, 13, 10, 14, 10, 182, 5, 10, 185,
	10, 10, 3, 10, 3, 10, 3, 10, 6, 10, 190, 10, 10, 13, 10, 14, 10, 191, 5,
	10, 194, 10, 10, 5, 10, 196, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 6, 11, 204, 10, 11, 13, 11, 14, 11, 205, 5, 11, 208, 10, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 6, 11, 215, 10, 11, 13, 11, 14, 11, 216,
	5, 11, 219, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 5, 11, 232, 10, 11, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 245, 10, 13, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 6, 15, 253, 10, 15, 13, 15, 14,
	15, 254, 3, 15, 5, 15, 258, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	264, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 272, 10,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 6, 18, 278, 10, 18, 13, 18, 14, 18, 279,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 6, 18, 287, 10, 18, 13, 18, 14, 18,
	288, 3, 18, 3, 18, 5, 18, 293, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19,
	299, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 5, 20, 323, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 7, 20, 331, 10, 20, 12, 20, 14, 20, 334, 11, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 5, 21, 340, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 7, 21, 348, 10, 21, 12, 21, 14, 21, 351, 11, 21, 3, 22, 3, 22, 3, 22,
	5, 22, 356, 10, 22, 3, 23, 3, 23, 5, 23, 360, 10, 23, 3, 24, 3, 24, 3,
	24, 5, 24, 365, 10, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 5, 26, 372,
	10, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 7, 26, 384, 10, 26, 12, 26, 14, 26, 387, 11, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 2, 5, 38, 40, 50, 28, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 2, 5, 3, 2,
	21, 22, 3, 2, 23, 24, 3, 2, 34, 35, 2, 439, 2, 82, 3, 2, 2, 2, 4, 86, 3,
	2, 2, 2, 6, 118, 3, 2, 2, 2, 8, 121, 3, 2, 2, 2, 10, 141, 3, 2, 2, 2, 12,
	150, 3, 2, 2, 2, 14, 152, 3, 2, 2, 2, 16, 172, 3, 2, 2, 2, 18, 195, 3,
	2, 2, 2, 20, 231, 3, 2, 2, 2, 22, 233, 3, 2, 2, 2, 24, 244, 3, 2, 2, 2,
	26, 246, 3, 2, 2, 2, 28, 257, 3, 2, 2, 2, 30, 263, 3, 2, 2, 2, 32, 271,
	3, 2, 2, 2, 34, 292, 3, 2, 2, 2, 36, 298, 3, 2, 2, 2, 38, 322, 3, 2, 2,
	2, 40, 339, 3, 2, 2, 2, 42, 355, 3, 2, 2, 2, 44, 359, 3, 2, 2, 2, 46, 364,
	3, 2, 2, 2, 48, 366, 3, 2, 2, 2, 50, 371, 3, 2, 2, 2, 52, 388, 3, 2, 2,
	2, 54, 61, 7, 6, 2, 2, 55, 57, 7, 6, 2, 2, 56, 55, 3, 2, 2, 2, 57, 58,
	3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 3, 2, 2, 2,
	60, 54, 3, 2, 2, 2, 60, 56, 3, 2, 2, 2, 61, 68, 3, 2, 2, 2, 62, 69, 5,
	4, 3, 2, 63, 65, 5, 4, 3, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66,
	64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 62, 3, 2, 2,
	2, 68, 64, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 71, 7, 2, 2, 3, 71, 83,
	3, 2, 2, 2, 72, 79, 5, 4, 3, 2, 73, 75, 5, 4, 3, 2, 74, 73, 3, 2, 2, 2,
	75, 76, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 79, 3,
	2, 2, 2, 78, 72, 3, 2, 2, 2, 78, 74, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80,
	81, 7, 2, 2, 3, 81, 83, 3, 2, 2, 2, 82, 60, 3, 2, 2, 2, 82, 78, 3, 2, 2,
	2, 83, 3, 3, 2, 2, 2, 84, 87, 5, 6, 4, 2, 85, 87, 5, 10, 6, 2, 86, 84,
	3, 2, 2, 2, 86, 85, 3, 2, 2, 2, 87, 5, 3, 2, 2, 2, 88, 89, 7, 31, 2, 2,
	89, 90, 5, 52, 27, 2, 90, 91, 7, 29, 2, 2, 91, 92, 7, 30, 2, 2, 92, 119,
	3, 2, 2, 2, 93, 95, 7, 31, 2, 2, 94, 96, 5, 52, 27, 2, 95, 94, 3, 2, 2,
	2, 96, 97, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99,
	3, 2, 2, 2, 99, 100, 7, 29, 2, 2, 100, 101, 7, 30, 2, 2, 101, 119, 3, 2,
	2, 2, 102, 103, 7, 31, 2, 2, 103, 104, 5, 52, 27, 2, 104, 105, 7, 29, 2,
	2, 105, 106, 5, 8, 5, 2, 106, 107, 7, 30, 2, 2, 107, 119, 3, 2, 2, 2, 108,
	110, 7, 31, 2, 2, 109, 111, 5, 52, 27, 2, 110, 109, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 3, 2,
	2, 2, 114, 115, 7, 29, 2, 2, 115, 116, 5, 8, 5, 2, 116, 117, 7, 30, 2,
	2, 117, 119, 3, 2, 2, 2, 118, 88, 3, 2, 2, 2, 118, 93, 3, 2, 2, 2, 118,
	102, 3, 2, 2, 2, 118, 108, 3, 2, 2, 2, 119, 7, 3, 2, 2, 2, 120, 122, 5,
	10, 6, 2, 121, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 121, 3, 2, 2,
	2, 123, 124, 3, 2, 2, 2, 124, 9, 3, 2, 2, 2, 125, 142, 5, 22, 12, 2, 126,
	142, 5, 18, 10, 2, 127, 142, 5, 14, 8, 2, 128, 129, 5, 12, 7, 2, 129, 130,
	7, 6, 2, 2, 130, 142, 3, 2, 2, 2, 131, 132, 5, 28, 15, 2, 132, 133, 7,
	6, 2, 2, 133, 142, 3, 2, 2, 2, 134, 135, 5, 16, 9, 2, 135, 136, 7, 6, 2,
	2, 136, 142, 3, 2, 2, 2, 137, 138, 5, 50, 26, 2, 138, 139, 7, 6, 2, 2,
	139, 142, 3, 2, 2, 2, 140, 142, 7, 6, 2, 2, 141, 125, 3, 2, 2, 2, 141,
	126, 3, 2, 2, 2, 141, 127, 3, 2, 2, 2, 141, 128, 3, 2, 2, 2, 141, 131,
	3, 2, 2, 2, 141, 134, 3, 2, 2, 2, 141, 137, 3, 2, 2, 2, 141, 140, 3, 2,
	2, 2, 142, 11, 3, 2, 2, 2, 143, 144, 7, 25, 2, 2, 144, 151, 5, 46, 24,
	2, 145, 146, 7, 25, 2, 2, 146, 147, 5, 46, 24, 2, 147, 148, 7, 27, 2, 2,
	148, 149, 5, 52, 27, 2, 149, 151, 3, 2, 2, 2, 150, 143, 3, 2, 2, 2, 150,
	145, 3, 2, 2, 2, 151, 13, 3, 2, 2, 2, 152, 155, 7, 26, 2, 2, 153, 156,
	5, 32, 17, 2, 154, 156, 5, 52, 27, 2, 155, 153, 3, 2, 2, 2, 155, 154, 3,
	2, 2, 2, 156, 15, 3, 2, 2, 2, 157, 158, 7, 28, 2, 2, 158, 159, 5, 38, 20,
	2, 159, 161, 7, 29, 2, 2, 160, 162, 5, 10, 6, 2, 161, 160, 3, 2, 2, 2,
	162, 163, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164,
	165, 3, 2, 2, 2, 165, 166, 7, 30, 2, 2, 166, 173, 3, 2, 2, 2, 167, 168,
	7, 28, 2, 2, 168, 169, 5, 38, 20, 2, 169, 170, 7, 29, 2, 2, 170, 171, 7,
	30, 2, 2, 171, 173, 3, 2, 2, 2, 172, 157, 3, 2, 2, 2, 172, 167, 3, 2, 2,
	2, 173, 17, 3, 2, 2, 2, 174, 175, 7, 32, 2, 2, 175, 196, 5, 32, 17, 2,
	176, 177, 7, 32, 2, 2, 177, 184, 5, 32, 17, 2, 178, 185, 7, 6, 2, 2, 179,
	181, 7, 6, 2, 2, 180, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 180,
	3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 185, 3, 2, 2, 2, 184, 178, 3, 2,
	2, 2, 184, 180, 3, 2, 2, 2, 185, 196, 3, 2, 2, 2, 186, 193, 7, 32, 2, 2,
	187, 194, 7, 6, 2, 2, 188, 190, 7, 6, 2, 2, 189, 188, 3, 2, 2, 2, 190,
	191, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 194,
	3, 2, 2, 2, 193, 187, 3, 2, 2, 2, 193, 189, 3, 2, 2, 2, 194, 196, 3, 2,
	2, 2, 195, 174, 3, 2, 2, 2, 195, 176, 3, 2, 2, 2, 195, 186, 3, 2, 2, 2,
	196, 19, 3, 2, 2, 2, 197, 198, 7, 17, 2, 2, 198, 232, 7, 18, 2, 2, 199,
	200, 7, 17, 2, 2, 200, 207, 7, 6, 2, 2, 201, 208, 5, 22, 12, 2, 202, 204,
	5, 22, 12, 2, 203, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 203, 3,
	2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 208, 3, 2, 2, 2, 207, 201, 3, 2, 2,
	2, 207, 203, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 210, 7, 18, 2, 2, 210,
	232, 3, 2, 2, 2, 211, 218, 7, 17, 2, 2, 212, 219, 5, 22, 12, 2, 213, 215,
	5, 22, 12, 2, 214, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 214, 3,
	2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 219, 3, 2, 2, 2, 218, 212, 3, 2, 2,
	2, 218, 214, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 7, 18, 2, 2, 221,
	232, 3, 2, 2, 2, 222, 223, 7, 17, 2, 2, 223, 224, 5, 24, 13, 2, 224, 225,
	7, 18, 2, 2, 225, 232, 3, 2, 2, 2, 226, 227, 7, 17, 2, 2, 227, 228, 7,
	6, 2, 2, 228, 229, 5, 24, 13, 2, 229, 230, 7, 18, 2, 2, 230, 232, 3, 2,
	2, 2, 231, 197, 3, 2, 2, 2, 231, 199, 3, 2, 2, 2, 231, 211, 3, 2, 2, 2,
	231, 222, 3, 2, 2, 2, 231, 226, 3, 2, 2, 2, 232, 21, 3, 2, 2, 2, 233, 234,
	5, 24, 13, 2, 234, 235, 7, 6, 2, 2, 235, 23, 3, 2, 2, 2, 236, 237, 5, 52,
	27, 2, 237, 238, 7, 9, 2, 2, 238, 239, 5, 32, 17, 2, 239, 245, 3, 2, 2,
	2, 240, 241, 5, 52, 27, 2, 241, 242, 7, 9, 2, 2, 242, 243, 5, 28, 15, 2,
	243, 245, 3, 2, 2, 2, 244, 236, 3, 2, 2, 2, 244, 240, 3, 2, 2, 2, 245,
	25, 3, 2, 2, 2, 246, 247, 7, 15, 2, 2, 247, 248, 5, 28, 15, 2, 248, 249,
	7, 16, 2, 2, 249, 27, 3, 2, 2, 2, 250, 252, 5, 52, 27, 2, 251, 253, 5,
	30, 16, 2, 252, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 252, 3, 2,
	2, 2, 254, 255, 3, 2, 2, 2, 255, 258, 3, 2, 2, 2, 256, 258, 5, 52, 27,
	2, 257, 250, 3, 2, 2, 2, 257, 256, 3, 2, 2, 2, 258, 29, 3, 2, 2, 2, 259,
	264, 7, 4, 2, 2, 260, 264, 7, 24, 2, 2, 261, 264, 7, 23, 2, 2, 262, 264,
	5, 32, 17, 2, 263, 259, 3, 2, 2, 2, 263, 260, 3, 2, 2, 2, 263, 261, 3,
	2, 2, 2, 263, 262, 3, 2, 2, 2, 264, 31, 3, 2, 2, 2, 265, 272, 5, 52, 27,
	2, 266, 272, 5, 42, 22, 2, 267, 272, 5, 40, 21, 2, 268, 272, 5, 26, 14,
	2, 269, 272, 5, 20, 11, 2, 270, 272, 5, 34, 18, 2, 271, 265, 3, 2, 2, 2,
	271, 266, 3, 2, 2, 2, 271, 267, 3, 2, 2, 2, 271, 268, 3, 2, 2, 2, 271,
	269, 3, 2, 2, 2, 271, 270, 3, 2, 2, 2, 272, 33, 3, 2, 2, 2, 273, 274, 7,
	19, 2, 2, 274, 293, 7, 20, 2, 2, 275, 277, 7, 19, 2, 2, 276, 278, 5, 36,
	19, 2, 277, 276, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2,
	279, 280, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 282, 7, 20, 2, 2, 282,
	293, 3, 2, 2, 2, 283, 284, 7, 19, 2, 2, 284, 286, 7, 6, 2, 2, 285, 287,
	5, 36, 19, 2, 286, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 286, 3,
	2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 291, 7, 20, 2,
	2, 291, 293, 3, 2, 2, 2, 292, 273, 3, 2, 2, 2, 292, 275, 3, 2, 2, 2, 292,
	283, 3, 2, 2, 2, 293, 35, 3, 2, 2, 2, 294, 299, 5, 32, 17, 2, 295, 296,
	5, 32, 17, 2, 296, 297, 7, 6, 2, 2, 297, 299, 3, 2, 2, 2, 298, 294, 3,
	2, 2, 2, 298, 295, 3, 2, 2, 2, 299, 37, 3, 2, 2, 2, 300, 301, 8, 20, 1,
	2, 301, 302, 5, 32, 17, 2, 302, 303, 7, 8, 2, 2, 303, 304, 5, 32, 17, 2,
	304, 323, 3, 2, 2, 2, 305, 306, 5, 32, 17, 2, 306, 307, 7, 11, 2, 2, 307,
	308, 5, 32, 17, 2, 308, 323, 3, 2, 2, 2, 309, 310, 5, 32, 17, 2, 310, 311,
	7, 12, 2, 2, 311, 312, 5, 32, 17, 2, 312, 323, 3, 2, 2, 2, 313, 314, 5,
	32, 17, 2, 314, 315, 7, 13, 2, 2, 315, 316, 5, 32, 17, 2, 316, 323, 3,
	2, 2, 2, 317, 318, 5, 32, 17, 2, 318, 319, 7, 14, 2, 2, 319, 320, 5, 32,
	17, 2, 320, 323, 3, 2, 2, 2, 321, 323, 5, 48, 25, 2, 322, 300, 3, 2, 2,
	2, 322, 305, 3, 2, 2, 2, 322, 309, 3, 2, 2, 2, 322, 313, 3, 2, 2, 2, 322,
	317, 3, 2, 2, 2, 322, 321, 3, 2, 2, 2, 323, 332, 3, 2, 2, 2, 324, 325,
	12, 5, 2, 2, 325, 326, 7, 36, 2, 2, 326, 331, 5, 38, 20, 6, 327, 328, 12,
	4, 2, 2, 328, 329, 7, 37, 2, 2, 329, 331, 5, 38, 20, 5, 330, 324, 3, 2,
	2, 2, 330, 327, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2,
	332, 333, 3, 2, 2, 2, 333, 39, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 335, 336,
	8, 21, 1, 2, 336, 337, 7, 22, 2, 2, 337, 340, 5, 40, 21, 4, 338, 340, 5,
	44, 23, 2, 339, 335, 3, 2, 2, 2, 339, 338, 3, 2, 2, 2, 340, 349, 3, 2,
	2, 2, 341, 342, 12, 6, 2, 2, 342, 343, 9, 2, 2, 2, 343, 348, 5, 40, 21,
	7, 344, 345, 12, 5, 2, 2, 345, 346, 9, 3, 2, 2, 346, 348, 5, 40, 21, 6,
	347, 341, 3, 2, 2, 2, 347, 344, 3, 2, 2, 2, 348, 351, 3, 2, 2, 2, 349,
	347, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350, 41, 3, 2, 2, 2, 351, 349, 3,
	2, 2, 2, 352, 356, 5, 44, 23, 2, 353, 356, 5, 46, 24, 2, 354, 356, 5, 48,
	25, 2, 355, 352, 3, 2, 2, 2, 355, 353, 3, 2, 2, 2, 355, 354, 3, 2, 2, 2,
	356, 43, 3, 2, 2, 2, 357, 360, 7, 38, 2, 2, 358, 360, 7, 39, 2, 2, 359,
	357, 3, 2, 2, 2, 359, 358, 3, 2, 2, 2, 360, 45, 3, 2, 2, 2, 361, 365, 7,
	41, 2, 2, 362, 365, 7, 42, 2, 2, 363, 365, 7, 40, 2, 2, 364, 361, 3, 2,
	2, 2, 364, 362, 3, 2, 2, 2, 364, 363, 3, 2, 2, 2, 365, 47, 3, 2, 2, 2,
	366, 367, 9, 4, 2, 2, 367, 49, 3, 2, 2, 2, 368, 369, 8, 26, 1, 2, 369,
	372, 5, 28, 15, 2, 370, 372, 5, 32, 17, 2, 371, 368, 3, 2, 2, 2, 371, 370,
	3, 2, 2, 2, 372, 385, 3, 2, 2, 2, 373, 374, 12, 6, 2, 2, 374, 375, 7, 10,
	2, 2, 375, 384, 5, 50, 26, 7, 376, 377, 12, 5, 2, 2, 377, 378, 7, 6, 2,
	2, 378, 379, 7, 10, 2, 2, 379, 384, 5, 50, 26, 6, 380, 381, 12, 7, 2, 2,
	381, 382, 7, 10, 2, 2, 382, 384, 5, 28, 15, 2, 383, 373, 3, 2, 2, 2, 383,
	376, 3, 2, 2, 2, 383, 380, 3, 2, 2, 2, 384, 387, 3, 2, 2, 2, 385, 383,
	3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 51, 3, 2, 2, 2, 387, 385, 3, 2,
	2, 2, 388, 389, 7, 3, 2, 2, 389, 390, 7, 41, 2, 2, 390, 53, 3, 2, 2, 2,
	50, 58, 60, 66, 68, 76, 78, 82, 86, 97, 112, 118, 123, 141, 150, 155, 163,
	172, 182, 184, 191, 193, 195, 205, 207, 216, 218, 231, 244, 254, 257, 263,
	271, 279, 288, 292, 298, 322, 330, 332, 339, 347, 349, 355, 359, 364, 371,
	383, 385,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "", "", "", "", "'=='", "'='", "'|'", "'<'", "'>'", "'<='",
	"'>='", "'('", "')'", "'{'", "'}'", "'['", "']'", "'+'", "'-'", "'/'",
	"'*'", "'import'", "'export'", "'as'", "'if'", "'do'", "'end'", "'func'",
	"'return'", "'for'", "'true'", "'false'", "'or'", "'and'",
}
var symbolicNames = []string{
	"", "", "FLAG", "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN",
	"PIPE", "LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN",
	"LSQUIG", "RSQUIG", "LBLOCK", "RBLOCK", "PLUS", "SUB", "DIV", "MULT", "IMPORT",
	"EXPORT", "AS", "IF", "DO", "END", "FUNC", "RETURN", "FOR", "TRUE", "FALSE",
	"OR", "AND", "INT", "FLOAT", "QUOTED", "LETTERS", "STRING",
}

var ruleNames = []string{
	"boop", "code", "funcdef", "funcguts", "statement", "importstat", "exportstat",
	"ifstat", "returnstat", "structexpr", "assignstat", "assign", "paren_fncall",
	"fncall", "fnargs", "term", "list", "listterm", "conditional", "math",
	"literal", "num", "words", "boolexpr", "pipe", "label",
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
	BeepBoopParserFLAG         = 2
	BeepBoopParserCOMMENT      = 3
	BeepBoopParserNEWLINE      = 4
	BeepBoopParserWHITESPACE   = 5
	BeepBoopParserEQUALS       = 6
	BeepBoopParserASSIGN       = 7
	BeepBoopParserPIPE         = 8
	BeepBoopParserLTHAN        = 9
	BeepBoopParserGTHAN        = 10
	BeepBoopParserLTHAN_EQUALS = 11
	BeepBoopParserGTHAN_EQUALS = 12
	BeepBoopParserLPAREN       = 13
	BeepBoopParserRPAREN       = 14
	BeepBoopParserLSQUIG       = 15
	BeepBoopParserRSQUIG       = 16
	BeepBoopParserLBLOCK       = 17
	BeepBoopParserRBLOCK       = 18
	BeepBoopParserPLUS         = 19
	BeepBoopParserSUB          = 20
	BeepBoopParserDIV          = 21
	BeepBoopParserMULT         = 22
	BeepBoopParserIMPORT       = 23
	BeepBoopParserEXPORT       = 24
	BeepBoopParserAS           = 25
	BeepBoopParserIF           = 26
	BeepBoopParserDO           = 27
	BeepBoopParserEND          = 28
	BeepBoopParserFUNC         = 29
	BeepBoopParserRETURN       = 30
	BeepBoopParserFOR          = 31
	BeepBoopParserTRUE         = 32
	BeepBoopParserFALSE        = 33
	BeepBoopParserOR           = 34
	BeepBoopParserAND          = 35
	BeepBoopParserINT          = 36
	BeepBoopParserFLOAT        = 37
	BeepBoopParserQUOTED       = 38
	BeepBoopParserLETTERS      = 39
	BeepBoopParserSTRING       = 40
)

// BeepBoopParser rules.
const (
	BeepBoopParserRULE_boop         = 0
	BeepBoopParserRULE_code         = 1
	BeepBoopParserRULE_funcdef      = 2
	BeepBoopParserRULE_funcguts     = 3
	BeepBoopParserRULE_statement    = 4
	BeepBoopParserRULE_importstat   = 5
	BeepBoopParserRULE_exportstat   = 6
	BeepBoopParserRULE_ifstat       = 7
	BeepBoopParserRULE_returnstat   = 8
	BeepBoopParserRULE_structexpr   = 9
	BeepBoopParserRULE_assignstat   = 10
	BeepBoopParserRULE_assign       = 11
	BeepBoopParserRULE_paren_fncall = 12
	BeepBoopParserRULE_fncall       = 13
	BeepBoopParserRULE_fnargs       = 14
	BeepBoopParserRULE_term         = 15
	BeepBoopParserRULE_list         = 16
	BeepBoopParserRULE_listterm     = 17
	BeepBoopParserRULE_conditional  = 18
	BeepBoopParserRULE_math         = 19
	BeepBoopParserRULE_literal      = 20
	BeepBoopParserRULE_num          = 21
	BeepBoopParserRULE_words        = 22
	BeepBoopParserRULE_boolexpr     = 23
	BeepBoopParserRULE_pipe         = 24
	BeepBoopParserRULE_label        = 25
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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(52)
				p.Match(BeepBoopParserNEWLINE)
			}

		case 2:
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(53)
						p.Match(BeepBoopParserNEWLINE)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(56)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(60)
				p.Code()
			}

		case 2:
			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserIMPORT)|(1<<BeepBoopParserEXPORT)|(1<<BeepBoopParserIF)|(1<<BeepBoopParserFUNC)|(1<<BeepBoopParserRETURN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserTRUE-32))|(1<<(BeepBoopParserFALSE-32))|(1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserFLOAT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32))|(1<<(BeepBoopParserSTRING-32)))) != 0) {
				{
					p.SetState(61)
					p.Code()
				}

				p.SetState(64)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(68)
			p.Match(BeepBoopParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(70)
				p.Code()
			}

		case 2:
			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserIMPORT)|(1<<BeepBoopParserEXPORT)|(1<<BeepBoopParserIF)|(1<<BeepBoopParserFUNC)|(1<<BeepBoopParserRETURN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserTRUE-32))|(1<<(BeepBoopParserFALSE-32))|(1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserFLOAT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32))|(1<<(BeepBoopParserSTRING-32)))) != 0) {
				{
					p.SetState(71)
					p.Code()
				}

				p.SetState(74)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(78)
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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserFUNC:
		localctx = NewFuncdefCodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Funcdef()
		}

	case BeepBoopParserT__0, BeepBoopParserNEWLINE, BeepBoopParserLPAREN, BeepBoopParserLSQUIG, BeepBoopParserLBLOCK, BeepBoopParserSUB, BeepBoopParserIMPORT, BeepBoopParserEXPORT, BeepBoopParserIF, BeepBoopParserRETURN, BeepBoopParserTRUE, BeepBoopParserFALSE, BeepBoopParserINT, BeepBoopParserFLOAT, BeepBoopParserQUOTED, BeepBoopParserLETTERS, BeepBoopParserSTRING:
		localctx = NewStatementCodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
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

	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(BeepBoopParserFUNC)
		}
		{
			p.SetState(87)
			p.Label()
		}
		{
			p.SetState(88)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(89)
			p.Match(BeepBoopParserEND)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Match(BeepBoopParserFUNC)
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == BeepBoopParserT__0 {
			{
				p.SetState(92)
				p.Label()
			}

			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(98)
			p.Match(BeepBoopParserEND)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Match(BeepBoopParserFUNC)
		}
		{
			p.SetState(101)
			p.Label()
		}
		{
			p.SetState(102)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(103)
			p.Funcguts()
		}
		{
			p.SetState(104)
			p.Match(BeepBoopParserEND)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.Match(BeepBoopParserFUNC)
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == BeepBoopParserT__0 {
			{
				p.SetState(107)
				p.Label()
			}

			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(112)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(113)
			p.Funcguts()
		}
		{
			p.SetState(114)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserIMPORT)|(1<<BeepBoopParserEXPORT)|(1<<BeepBoopParserIF)|(1<<BeepBoopParserRETURN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserTRUE-32))|(1<<(BeepBoopParserFALSE-32))|(1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserFLOAT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32))|(1<<(BeepBoopParserSTRING-32)))) != 0) {
		{
			p.SetState(118)
			p.Statement()
		}

		p.SetState(121)
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

type ExportStatementContext struct {
	*StatementContext
}

func NewExportStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExportStatementContext {
	var p = new(ExportStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ExportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportStatementContext) Exportstat() IExportstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportstatContext)
}

func (s *ExportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterExportStatement(s)
	}
}

func (s *ExportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitExportStatement(s)
	}
}

func (s *ExportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitExportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImportStatementContext struct {
	*StatementContext
}

func NewImportStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) Importstat() IImportstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportstatContext)
}

func (s *ImportStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitImportStatement(s)

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

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Assignstat()
		}

	case 2:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Returnstat()
		}

	case 3:
		localctx = NewExportStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Exportstat()
		}

	case 4:
		localctx = NewImportStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.Importstat()
		}
		{
			p.SetState(127)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 5:
		localctx = NewFncallStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(129)
			p.Fncall()
		}
		{
			p.SetState(130)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 6:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
			p.Ifstat()
		}
		{
			p.SetState(133)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 7:
		localctx = NewPipeStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(135)
			p.pipe(0)
		}
		{
			p.SetState(136)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 8:
		localctx = NewNewlineStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(138)
			p.Match(BeepBoopParserNEWLINE)
		}

	}

	return localctx
}

// IImportstatContext is an interface to support dynamic dispatch.
type IImportstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportstatContext differentiates from other interfaces.
	IsImportstatContext()
}

type ImportstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportstatContext() *ImportstatContext {
	var p = new(ImportstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_importstat
	return p
}

func (*ImportstatContext) IsImportstatContext() {}

func NewImportstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportstatContext {
	var p = new(ImportstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_importstat

	return p
}

func (s *ImportstatContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportstatContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserIMPORT, 0)
}

func (s *ImportstatContext) Words() IWordsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWordsContext)
}

func (s *ImportstatContext) AS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserAS, 0)
}

func (s *ImportstatContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *ImportstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterImportstat(s)
	}
}

func (s *ImportstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitImportstat(s)
	}
}

func (s *ImportstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitImportstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Importstat() (localctx IImportstatContext) {
	localctx = NewImportstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BeepBoopParserRULE_importstat)

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

	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(BeepBoopParserIMPORT)
		}
		{
			p.SetState(142)
			p.Words()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.Match(BeepBoopParserIMPORT)
		}
		{
			p.SetState(144)
			p.Words()
		}
		{
			p.SetState(145)
			p.Match(BeepBoopParserAS)
		}
		{
			p.SetState(146)
			p.Label()
		}

	}

	return localctx
}

// IExportstatContext is an interface to support dynamic dispatch.
type IExportstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportstatContext differentiates from other interfaces.
	IsExportstatContext()
}

type ExportstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportstatContext() *ExportstatContext {
	var p = new(ExportstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_exportstat
	return p
}

func (*ExportstatContext) IsExportstatContext() {}

func NewExportstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportstatContext {
	var p = new(ExportstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_exportstat

	return p
}

func (s *ExportstatContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportstatContext) EXPORT() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEXPORT, 0)
}

func (s *ExportstatContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExportstatContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *ExportstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterExportstat(s)
	}
}

func (s *ExportstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitExportstat(s)
	}
}

func (s *ExportstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitExportstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Exportstat() (localctx IExportstatContext) {
	localctx = NewExportstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BeepBoopParserRULE_exportstat)

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
		p.SetState(150)
		p.Match(BeepBoopParserEXPORT)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(151)
			p.Term()
		}

	case 2:
		{
			p.SetState(152)
			p.Label()
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
	p.EnterRule(localctx, 14, BeepBoopParserRULE_ifstat)
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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Match(BeepBoopParserIF)
		}
		{
			p.SetState(156)
			p.conditional(0)
		}
		{
			p.SetState(157)
			p.Match(BeepBoopParserDO)
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserSUB)|(1<<BeepBoopParserIMPORT)|(1<<BeepBoopParserEXPORT)|(1<<BeepBoopParserIF)|(1<<BeepBoopParserRETURN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserTRUE-32))|(1<<(BeepBoopParserFALSE-32))|(1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserFLOAT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32))|(1<<(BeepBoopParserSTRING-32)))) != 0) {
			{
				p.SetState(158)
				p.Statement()
			}

			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(163)
			p.Match(BeepBoopParserEND)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(BeepBoopParserIF)
		}
		{
			p.SetState(166)
			p.conditional(0)
		}
		{
			p.SetState(167)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(168)
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
	p.EnterRule(localctx, 16, BeepBoopParserRULE_returnstat)

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

	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Match(BeepBoopParserRETURN)
		}
		{
			p.SetState(173)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Match(BeepBoopParserRETURN)
		}
		{
			p.SetState(175)
			p.Term()
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(176)
				p.Match(BeepBoopParserNEWLINE)
			}

		case 2:
			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(177)
						p.Match(BeepBoopParserNEWLINE)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(180)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(184)
			p.Match(BeepBoopParserRETURN)
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(185)
				p.Match(BeepBoopParserNEWLINE)
			}

		case 2:
			p.SetState(187)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(186)
						p.Match(BeepBoopParserNEWLINE)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(189)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 18, BeepBoopParserRULE_structexpr)
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.Match(BeepBoopParserLSQUIG)
		}
		{
			p.SetState(196)
			p.Match(BeepBoopParserRSQUIG)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Match(BeepBoopParserLSQUIG)
		}
		{
			p.SetState(198)
			p.Match(BeepBoopParserNEWLINE)
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(199)
				p.Assignstat()
			}

		case 2:
			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == BeepBoopParserT__0 {
				{
					p.SetState(200)
					p.Assignstat()
				}

				p.SetState(203)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(207)
			p.Match(BeepBoopParserRSQUIG)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(209)
			p.Match(BeepBoopParserLSQUIG)
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(210)
				p.Assignstat()
			}

		case 2:
			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == BeepBoopParserT__0 {
				{
					p.SetState(211)
					p.Assignstat()
				}

				p.SetState(214)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(218)
			p.Match(BeepBoopParserRSQUIG)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
			p.Match(BeepBoopParserLSQUIG)
		}
		{
			p.SetState(221)
			p.Assign()
		}
		{
			p.SetState(222)
			p.Match(BeepBoopParserRSQUIG)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(224)
			p.Match(BeepBoopParserLSQUIG)
		}
		{
			p.SetState(225)
			p.Match(BeepBoopParserNEWLINE)
		}
		{
			p.SetState(226)
			p.Assign()
		}
		{
			p.SetState(227)
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
	p.EnterRule(localctx, 20, BeepBoopParserRULE_assignstat)

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
		p.SetState(231)
		p.Assign()
	}
	{
		p.SetState(232)
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
	p.EnterRule(localctx, 22, BeepBoopParserRULE_assign)

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

	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Label()
		}
		{
			p.SetState(235)
			p.Match(BeepBoopParserASSIGN)
		}
		{
			p.SetState(236)
			p.Term()
		}

	case 2:
		localctx = NewFncallAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Label()
		}
		{
			p.SetState(239)
			p.Match(BeepBoopParserASSIGN)
		}
		{
			p.SetState(240)
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
	p.EnterRule(localctx, 24, BeepBoopParserRULE_paren_fncall)

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
		p.SetState(244)
		p.Match(BeepBoopParserLPAREN)
	}
	{
		p.SetState(245)
		p.Fncall()
	}
	{
		p.SetState(246)
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

func (s *FncallContext) AllFnargs() []IFnargsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFnargsContext)(nil)).Elem())
	var tst = make([]IFnargsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFnargsContext)
		}
	}

	return tst
}

func (s *FncallContext) Fnargs(i int) IFnargsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnargsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFnargsContext)
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
	p.EnterRule(localctx, 26, BeepBoopParserRULE_fncall)

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

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Label()
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(249)
					p.Fnargs()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Label()
		}

	}

	return localctx
}

// IFnargsContext is an interface to support dynamic dispatch.
type IFnargsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnargsContext differentiates from other interfaces.
	IsFnargsContext()
}

type FnargsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnargsContext() *FnargsContext {
	var p = new(FnargsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_fnargs
	return p
}

func (*FnargsContext) IsFnargsContext() {}

func NewFnargsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnargsContext {
	var p = new(FnargsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_fnargs

	return p
}

func (s *FnargsContext) GetParser() antlr.Parser { return s.parser }

func (s *FnargsContext) CopyFrom(ctx *FnargsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FnargsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnargsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TermFnargsContext struct {
	*FnargsContext
}

func NewTermFnargsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermFnargsContext {
	var p = new(TermFnargsContext)

	p.FnargsContext = NewEmptyFnargsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FnargsContext))

	return p
}

func (s *TermFnargsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermFnargsContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermFnargsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterTermFnargs(s)
	}
}

func (s *TermFnargsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitTermFnargs(s)
	}
}

func (s *TermFnargsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitTermFnargs(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivFnargsContext struct {
	*FnargsContext
}

func NewDivFnargsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivFnargsContext {
	var p = new(DivFnargsContext)

	p.FnargsContext = NewEmptyFnargsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FnargsContext))

	return p
}

func (s *DivFnargsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivFnargsContext) DIV() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserDIV, 0)
}

func (s *DivFnargsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterDivFnargs(s)
	}
}

func (s *DivFnargsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitDivFnargs(s)
	}
}

func (s *DivFnargsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitDivFnargs(s)

	default:
		return t.VisitChildren(s)
	}
}

type FlagFnargsContext struct {
	*FnargsContext
}

func NewFlagFnargsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FlagFnargsContext {
	var p = new(FlagFnargsContext)

	p.FnargsContext = NewEmptyFnargsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FnargsContext))

	return p
}

func (s *FlagFnargsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlagFnargsContext) FLAG() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserFLAG, 0)
}

func (s *FlagFnargsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFlagFnargs(s)
	}
}

func (s *FlagFnargsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFlagFnargs(s)
	}
}

func (s *FlagFnargsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFlagFnargs(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultFnargsContext struct {
	*FnargsContext
}

func NewMultFnargsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultFnargsContext {
	var p = new(MultFnargsContext)

	p.FnargsContext = NewEmptyFnargsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FnargsContext))

	return p
}

func (s *MultFnargsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultFnargsContext) MULT() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserMULT, 0)
}

func (s *MultFnargsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterMultFnargs(s)
	}
}

func (s *MultFnargsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitMultFnargs(s)
	}
}

func (s *MultFnargsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitMultFnargs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Fnargs() (localctx IFnargsContext) {
	localctx = NewFnargsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BeepBoopParserRULE_fnargs)

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

	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserFLAG:
		localctx = NewFlagFnargsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.Match(BeepBoopParserFLAG)
		}

	case BeepBoopParserMULT:
		localctx = NewMultFnargsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.Match(BeepBoopParserMULT)
		}

	case BeepBoopParserDIV:
		localctx = NewDivFnargsContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(259)
			p.Match(BeepBoopParserDIV)
		}

	case BeepBoopParserT__0, BeepBoopParserLPAREN, BeepBoopParserLSQUIG, BeepBoopParserLBLOCK, BeepBoopParserSUB, BeepBoopParserTRUE, BeepBoopParserFALSE, BeepBoopParserINT, BeepBoopParserFLOAT, BeepBoopParserQUOTED, BeepBoopParserLETTERS, BeepBoopParserSTRING:
		localctx = NewTermFnargsContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(260)
			p.Term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 30, BeepBoopParserRULE_term)

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

	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLabelTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Label()
		}

	case 2:
		localctx = NewLiteralTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.Literal()
		}

	case 3:
		localctx = NewMathTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(265)
			p.math(0)
		}

	case 4:
		localctx = NewParenfncallTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(266)
			p.Paren_fncall()
		}

	case 5:
		localctx = NewStructTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(267)
			p.Structexpr()
		}

	case 6:
		localctx = NewListTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(268)
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
	p.EnterRule(localctx, 32, BeepBoopParserRULE_list)
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

	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Match(BeepBoopParserLBLOCK)
		}
		{
			p.SetState(272)
			p.Match(BeepBoopParserRBLOCK)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.Match(BeepBoopParserLBLOCK)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserSUB))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserTRUE-32))|(1<<(BeepBoopParserFALSE-32))|(1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserFLOAT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32))|(1<<(BeepBoopParserSTRING-32)))) != 0) {
			{
				p.SetState(274)
				p.Listterm()
			}

			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(279)
			p.Match(BeepBoopParserRBLOCK)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(281)
			p.Match(BeepBoopParserLBLOCK)
		}
		{
			p.SetState(282)
			p.Match(BeepBoopParserNEWLINE)
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserLPAREN)|(1<<BeepBoopParserLSQUIG)|(1<<BeepBoopParserLBLOCK)|(1<<BeepBoopParserSUB))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BeepBoopParserTRUE-32))|(1<<(BeepBoopParserFALSE-32))|(1<<(BeepBoopParserINT-32))|(1<<(BeepBoopParserFLOAT-32))|(1<<(BeepBoopParserQUOTED-32))|(1<<(BeepBoopParserLETTERS-32))|(1<<(BeepBoopParserSTRING-32)))) != 0) {
			{
				p.SetState(283)
				p.Listterm()
			}

			p.SetState(286)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(288)
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
	p.EnterRule(localctx, 34, BeepBoopParserRULE_listterm)

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

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(292)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			p.Term()
		}
		{
			p.SetState(294)
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
	_startState := 36
	p.EnterRecursionRule(localctx, 36, BeepBoopParserRULE_conditional, _p)

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
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEqualsCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(299)
			p.Term()
		}
		{
			p.SetState(300)
			p.Match(BeepBoopParserEQUALS)
		}
		{
			p.SetState(301)
			p.Term()
		}

	case 2:
		localctx = NewLthanCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(303)
			p.Term()
		}
		{
			p.SetState(304)
			p.Match(BeepBoopParserLTHAN)
		}
		{
			p.SetState(305)
			p.Term()
		}

	case 3:
		localctx = NewGthanCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(307)
			p.Term()
		}
		{
			p.SetState(308)
			p.Match(BeepBoopParserGTHAN)
		}
		{
			p.SetState(309)
			p.Term()
		}

	case 4:
		localctx = NewLthanequalsCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(311)
			p.Term()
		}
		{
			p.SetState(312)
			p.Match(BeepBoopParserLTHAN_EQUALS)
		}
		{
			p.SetState(313)
			p.Term()
		}

	case 5:
		localctx = NewGthanequalsCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(315)
			p.Term()
		}
		{
			p.SetState(316)
			p.Match(BeepBoopParserGTHAN_EQUALS)
		}
		{
			p.SetState(317)
			p.Term()
		}

	case 6:
		localctx = NewBoolCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(319)
			p.Boolexpr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(328)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOrCondContext(p, NewConditionalContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_conditional)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(323)
					p.Match(BeepBoopParserOR)
				}
				{
					p.SetState(324)
					p.conditional(4)
				}

			case 2:
				localctx = NewAndCondContext(p, NewConditionalContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_conditional)
				p.SetState(325)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(326)
					p.Match(BeepBoopParserAND)
				}
				{
					p.SetState(327)
					p.conditional(3)
				}

			}

		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
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

type MultiplicativeMathContext struct {
	*MathContext
	op antlr.Token
}

func NewMultiplicativeMathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeMathContext {
	var p = new(MultiplicativeMathContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *MultiplicativeMathContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicativeMathContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicativeMathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeMathContext) AllMath() []IMathContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathContext)(nil)).Elem())
	var tst = make([]IMathContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathContext)
		}
	}

	return tst
}

func (s *MultiplicativeMathContext) Math(i int) IMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *MultiplicativeMathContext) MULT() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserMULT, 0)
}

func (s *MultiplicativeMathContext) DIV() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserDIV, 0)
}

func (s *MultiplicativeMathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterMultiplicativeMath(s)
	}
}

func (s *MultiplicativeMathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitMultiplicativeMath(s)
	}
}

func (s *MultiplicativeMathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitMultiplicativeMath(s)

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
	_startState := 38
	p.EnterRecursionRule(localctx, 38, BeepBoopParserRULE_math, _p)
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
	p.SetState(337)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserSUB:
		localctx = NewUnarySubMathContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(334)
			p.Match(BeepBoopParserSUB)
		}
		{
			p.SetState(335)
			p.math(2)
		}

	case BeepBoopParserINT, BeepBoopParserFLOAT:
		localctx = NewSoloMathContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(336)
			p.Num()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(345)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAdditiveMathContext(p, NewMathContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_math)
				p.SetState(339)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(340)

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
					p.SetState(341)
					p.math(5)
				}

			case 2:
				localctx = NewMultiplicativeMathContext(p, NewMathContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_math)
				p.SetState(342)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(343)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicativeMathContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == BeepBoopParserDIV || _la == BeepBoopParserMULT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicativeMathContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(344)
					p.math(4)
				}

			}

		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
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

type WordsLiteralContext struct {
	*LiteralContext
}

func NewWordsLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WordsLiteralContext {
	var p = new(WordsLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *WordsLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordsLiteralContext) Words() IWordsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWordsContext)
}

func (s *WordsLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterWordsLiteral(s)
	}
}

func (s *WordsLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitWordsLiteral(s)
	}
}

func (s *WordsLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitWordsLiteral(s)

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
	p.EnterRule(localctx, 40, BeepBoopParserRULE_literal)

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

	p.SetState(353)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserINT, BeepBoopParserFLOAT:
		localctx = NewNumLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
			p.Num()
		}

	case BeepBoopParserQUOTED, BeepBoopParserLETTERS, BeepBoopParserSTRING:
		localctx = NewWordsLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(351)
			p.Words()
		}

	case BeepBoopParserTRUE, BeepBoopParserFALSE:
		localctx = NewBoolLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(352)
			p.Boolexpr()
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

func (s *NumContext) CopyFrom(ctx *NumContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FloatNumContext struct {
	*NumContext
}

func NewFloatNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatNumContext {
	var p = new(FloatNumContext)

	p.NumContext = NewEmptyNumContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumContext))

	return p
}

func (s *FloatNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatNumContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserFLOAT, 0)
}

func (s *FloatNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFloatNum(s)
	}
}

func (s *FloatNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFloatNum(s)
	}
}

func (s *FloatNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFloatNum(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntNumContext struct {
	*NumContext
}

func NewIntNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntNumContext {
	var p = new(IntNumContext)

	p.NumContext = NewEmptyNumContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumContext))

	return p
}

func (s *IntNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntNumContext) INT() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserINT, 0)
}

func (s *IntNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterIntNum(s)
	}
}

func (s *IntNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitIntNum(s)
	}
}

func (s *IntNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitIntNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Num() (localctx INumContext) {
	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BeepBoopParserRULE_num)

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

	p.SetState(357)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserINT:
		localctx = NewIntNumContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.Match(BeepBoopParserINT)
		}

	case BeepBoopParserFLOAT:
		localctx = NewFloatNumContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.Match(BeepBoopParserFLOAT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWordsContext is an interface to support dynamic dispatch.
type IWordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordsContext differentiates from other interfaces.
	IsWordsContext()
}

type WordsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordsContext() *WordsContext {
	var p = new(WordsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_words
	return p
}

func (*WordsContext) IsWordsContext() {}

func NewWordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordsContext {
	var p = new(WordsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_words

	return p
}

func (s *WordsContext) GetParser() antlr.Parser { return s.parser }

func (s *WordsContext) CopyFrom(ctx *WordsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *WordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QuotedWordsContext struct {
	*WordsContext
}

func NewQuotedWordsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedWordsContext {
	var p = new(QuotedWordsContext)

	p.WordsContext = NewEmptyWordsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WordsContext))

	return p
}

func (s *QuotedWordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedWordsContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserQUOTED, 0)
}

func (s *QuotedWordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterQuotedWords(s)
	}
}

func (s *QuotedWordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitQuotedWords(s)
	}
}

func (s *QuotedWordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitQuotedWords(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetterWordsContext struct {
	*WordsContext
}

func NewLetterWordsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetterWordsContext {
	var p = new(LetterWordsContext)

	p.WordsContext = NewEmptyWordsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WordsContext))

	return p
}

func (s *LetterWordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetterWordsContext) LETTERS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserLETTERS, 0)
}

func (s *LetterWordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterLetterWords(s)
	}
}

func (s *LetterWordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitLetterWords(s)
	}
}

func (s *LetterWordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitLetterWords(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringWordsContext struct {
	*WordsContext
}

func NewStringWordsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringWordsContext {
	var p = new(StringWordsContext)

	p.WordsContext = NewEmptyWordsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WordsContext))

	return p
}

func (s *StringWordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringWordsContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserSTRING, 0)
}

func (s *StringWordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterStringWords(s)
	}
}

func (s *StringWordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitStringWords(s)
	}
}

func (s *StringWordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitStringWords(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Words() (localctx IWordsContext) {
	localctx = NewWordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BeepBoopParserRULE_words)

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

	p.SetState(362)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserLETTERS:
		localctx = NewLetterWordsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(359)
			p.Match(BeepBoopParserLETTERS)
		}

	case BeepBoopParserSTRING:
		localctx = NewStringWordsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(360)
			p.Match(BeepBoopParserSTRING)
		}

	case BeepBoopParserQUOTED:
		localctx = NewQuotedWordsContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(361)
			p.Match(BeepBoopParserQUOTED)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 46, BeepBoopParserRULE_boolexpr)
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
		p.SetState(364)
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

func (s *PipeContext) CopyFrom(ctx *PipeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PipeToFncallPipeContext struct {
	*PipeContext
}

func NewPipeToFncallPipeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PipeToFncallPipeContext {
	var p = new(PipeToFncallPipeContext)

	p.PipeContext = NewEmptyPipeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PipeContext))

	return p
}

func (s *PipeToFncallPipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeToFncallPipeContext) Pipe() IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *PipeToFncallPipeContext) PIPE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserPIPE, 0)
}

func (s *PipeToFncallPipeContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *PipeToFncallPipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterPipeToFncallPipe(s)
	}
}

func (s *PipeToFncallPipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitPipeToFncallPipe(s)
	}
}

func (s *PipeToFncallPipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitPipeToFncallPipe(s)

	default:
		return t.VisitChildren(s)
	}
}

type PipeToPipeContext struct {
	*PipeContext
}

func NewPipeToPipeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PipeToPipeContext {
	var p = new(PipeToPipeContext)

	p.PipeContext = NewEmptyPipeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PipeContext))

	return p
}

func (s *PipeToPipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeToPipeContext) AllPipe() []IPipeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipeContext)(nil)).Elem())
	var tst = make([]IPipeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipeContext)
		}
	}

	return tst
}

func (s *PipeToPipeContext) Pipe(i int) IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *PipeToPipeContext) PIPE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserPIPE, 0)
}

func (s *PipeToPipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterPipeToPipe(s)
	}
}

func (s *PipeToPipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitPipeToPipe(s)
	}
}

func (s *PipeToPipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitPipeToPipe(s)

	default:
		return t.VisitChildren(s)
	}
}

type TermPipeContext struct {
	*PipeContext
}

func NewTermPipeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermPipeContext {
	var p = new(TermPipeContext)

	p.PipeContext = NewEmptyPipeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PipeContext))

	return p
}

func (s *TermPipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermPipeContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermPipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterTermPipe(s)
	}
}

func (s *TermPipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitTermPipe(s)
	}
}

func (s *TermPipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitTermPipe(s)

	default:
		return t.VisitChildren(s)
	}
}

type NewlinePipeContext struct {
	*PipeContext
}

func NewNewlinePipeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewlinePipeContext {
	var p = new(NewlinePipeContext)

	p.PipeContext = NewEmptyPipeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PipeContext))

	return p
}

func (s *NewlinePipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewlinePipeContext) AllPipe() []IPipeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipeContext)(nil)).Elem())
	var tst = make([]IPipeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipeContext)
		}
	}

	return tst
}

func (s *NewlinePipeContext) Pipe(i int) IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *NewlinePipeContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *NewlinePipeContext) PIPE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserPIPE, 0)
}

func (s *NewlinePipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterNewlinePipe(s)
	}
}

func (s *NewlinePipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitNewlinePipe(s)
	}
}

func (s *NewlinePipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitNewlinePipe(s)

	default:
		return t.VisitChildren(s)
	}
}

type FncallPipeContext struct {
	*PipeContext
}

func NewFncallPipeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FncallPipeContext {
	var p = new(FncallPipeContext)

	p.PipeContext = NewEmptyPipeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PipeContext))

	return p
}

func (s *FncallPipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FncallPipeContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *FncallPipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFncallPipe(s)
	}
}

func (s *FncallPipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFncallPipe(s)
	}
}

func (s *FncallPipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFncallPipe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Pipe() (localctx IPipeContext) {
	return p.pipe(0)
}

func (p *BeepBoopParser) pipe(_p int) (localctx IPipeContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPipeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPipeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, BeepBoopParserRULE_pipe, _p)

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
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFncallPipeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(367)
			p.Fncall()
		}

	case 2:
		localctx = NewTermPipeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(368)
			p.Term()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(381)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPipeToPipeContext(p, NewPipeContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_pipe)
				p.SetState(371)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(372)
					p.Match(BeepBoopParserPIPE)
				}
				{
					p.SetState(373)
					p.pipe(5)
				}

			case 2:
				localctx = NewNewlinePipeContext(p, NewPipeContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_pipe)
				p.SetState(374)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(375)
					p.Match(BeepBoopParserNEWLINE)
				}
				{
					p.SetState(376)
					p.Match(BeepBoopParserPIPE)
				}
				{
					p.SetState(377)
					p.pipe(4)
				}

			case 3:
				localctx = NewPipeToFncallPipeContext(p, NewPipeContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_pipe)
				p.SetState(378)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(379)
					p.Match(BeepBoopParserPIPE)
				}
				{
					p.SetState(380)
					p.Fncall()
				}

			}

		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 50, BeepBoopParserRULE_label)

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
		p.SetState(386)
		p.Match(BeepBoopParserT__0)
	}
	{
		p.SetState(387)
		p.Match(BeepBoopParserLETTERS)
	}

	return localctx
}

func (p *BeepBoopParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ConditionalContext = nil
		if localctx != nil {
			t = localctx.(*ConditionalContext)
		}
		return p.Conditional_Sempred(t, predIndex)

	case 19:
		var t *MathContext = nil
		if localctx != nil {
			t = localctx.(*MathContext)
		}
		return p.Math_Sempred(t, predIndex)

	case 24:
		var t *PipeContext = nil
		if localctx != nil {
			t = localctx.(*PipeContext)
		}
		return p.Pipe_Sempred(t, predIndex)

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
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *BeepBoopParser) Pipe_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
