// Code generated from /Users/renyunyi/gengine/internal/iantlr/gengine.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 52, 467, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 3, 2, 3, 2, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 
	10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 
	3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 
	20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 
	3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 
	31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 5, 33, 234, 10, 33, 3, 33, 6, 33, 
	237, 10, 33, 13, 33, 14, 33, 238, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 
	35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 
	3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 
	40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 
	3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 
	44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 
	3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 
	47, 3, 47, 3, 48, 6, 48, 312, 10, 48, 13, 48, 14, 48, 313, 3, 48, 7, 48, 
	317, 10, 48, 12, 48, 14, 48, 320, 11, 48, 3, 49, 6, 49, 323, 10, 49, 13, 
	49, 14, 49, 324, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 
	3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 
	58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 
	3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 
	65, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 
	3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 3, 74, 3, 75, 3, 
	75, 3, 75, 3, 75, 3, 75, 3, 75, 7, 75, 392, 10, 75, 12, 75, 14, 75, 395, 
	11, 75, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 3, 76, 3, 77, 6, 77, 404, 10, 
	77, 13, 77, 14, 77, 405, 5, 77, 408, 10, 77, 3, 77, 3, 77, 6, 77, 412, 
	10, 77, 13, 77, 14, 77, 413, 3, 77, 6, 77, 417, 10, 77, 13, 77, 14, 77, 
	418, 3, 77, 3, 77, 3, 77, 3, 77, 6, 77, 425, 10, 77, 13, 77, 14, 77, 426, 
	5, 77, 429, 10, 77, 3, 77, 3, 77, 6, 77, 433, 10, 77, 13, 77, 14, 77, 434, 
	3, 77, 3, 77, 3, 77, 6, 77, 440, 10, 77, 13, 77, 14, 77, 441, 3, 77, 3, 
	77, 5, 77, 446, 10, 77, 3, 78, 3, 78, 3, 78, 3, 78, 7, 78, 452, 10, 78, 
	12, 78, 14, 78, 455, 11, 78, 3, 78, 3, 78, 3, 78, 3, 78, 3, 79, 6, 79, 
	462, 10, 79, 13, 79, 14, 79, 463, 3, 79, 3, 79, 3, 453, 2, 80, 3, 3, 5, 
	4, 7, 5, 9, 6, 11, 2, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 
	27, 2, 29, 2, 31, 2, 33, 2, 35, 2, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2, 47, 
	2, 49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 7, 
	69, 8, 71, 9, 73, 10, 75, 11, 77, 12, 79, 13, 81, 14, 83, 15, 85, 16, 87, 
	17, 89, 18, 91, 19, 93, 20, 95, 21, 97, 22, 99, 23, 101, 24, 103, 25, 105, 
	26, 107, 27, 109, 28, 111, 29, 113, 30, 115, 31, 117, 32, 119, 33, 121, 
	34, 123, 35, 125, 36, 127, 37, 129, 38, 131, 39, 133, 40, 135, 41, 137, 
	42, 139, 43, 141, 44, 143, 45, 145, 46, 147, 47, 149, 48, 151, 49, 153, 
	50, 155, 51, 157, 52, 3, 2, 33, 3, 2, 50, 59, 4, 2, 67, 67, 99, 99, 4, 
	2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 
	2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 
	2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 
	2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 
	2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 
	2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 
	2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 
	2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 
	2, 92, 92, 124, 124, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 
	97, 97, 99, 124, 4, 2, 36, 36, 94, 94, 5, 2, 11, 12, 15, 15, 34, 34, 2, 
	459, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 
	2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 
	2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 
	3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 
	89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 
	2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 
	2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 
	3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 
	2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 
	2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 
	133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 
	2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 
	3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 
	2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 3, 159, 3, 2, 2, 2, 5, 161, 3, 
	2, 2, 2, 7, 167, 3, 2, 2, 2, 9, 171, 3, 2, 2, 2, 11, 177, 3, 2, 2, 2, 13, 
	179, 3, 2, 2, 2, 15, 181, 3, 2, 2, 2, 17, 183, 3, 2, 2, 2, 19, 185, 3, 
	2, 2, 2, 21, 187, 3, 2, 2, 2, 23, 189, 3, 2, 2, 2, 25, 191, 3, 2, 2, 2, 
	27, 193, 3, 2, 2, 2, 29, 195, 3, 2, 2, 2, 31, 197, 3, 2, 2, 2, 33, 199, 
	3, 2, 2, 2, 35, 201, 3, 2, 2, 2, 37, 203, 3, 2, 2, 2, 39, 205, 3, 2, 2, 
	2, 41, 207, 3, 2, 2, 2, 43, 209, 3, 2, 2, 2, 45, 211, 3, 2, 2, 2, 47, 213, 
	3, 2, 2, 2, 49, 215, 3, 2, 2, 2, 51, 217, 3, 2, 2, 2, 53, 219, 3, 2, 2, 
	2, 55, 221, 3, 2, 2, 2, 57, 223, 3, 2, 2, 2, 59, 225, 3, 2, 2, 2, 61, 227, 
	3, 2, 2, 2, 63, 229, 3, 2, 2, 2, 65, 231, 3, 2, 2, 2, 67, 240, 3, 2, 2, 
	2, 69, 244, 3, 2, 2, 2, 71, 249, 3, 2, 2, 2, 73, 252, 3, 2, 2, 2, 75, 255, 
	3, 2, 2, 2, 77, 260, 3, 2, 2, 2, 79, 263, 3, 2, 2, 2, 81, 268, 3, 2, 2, 
	2, 83, 275, 3, 2, 2, 2, 85, 280, 3, 2, 2, 2, 87, 286, 3, 2, 2, 2, 89, 291, 
	3, 2, 2, 2, 91, 300, 3, 2, 2, 2, 93, 306, 3, 2, 2, 2, 95, 311, 3, 2, 2, 
	2, 97, 322, 3, 2, 2, 2, 99, 326, 3, 2, 2, 2, 101, 328, 3, 2, 2, 2, 103, 
	330, 3, 2, 2, 2, 105, 332, 3, 2, 2, 2, 107, 334, 3, 2, 2, 2, 109, 337, 
	3, 2, 2, 2, 111, 339, 3, 2, 2, 2, 113, 341, 3, 2, 2, 2, 115, 344, 3, 2, 
	2, 2, 117, 347, 3, 2, 2, 2, 119, 350, 3, 2, 2, 2, 121, 352, 3, 2, 2, 2, 
	123, 355, 3, 2, 2, 2, 125, 357, 3, 2, 2, 2, 127, 360, 3, 2, 2, 2, 129, 
	363, 3, 2, 2, 2, 131, 366, 3, 2, 2, 2, 133, 369, 3, 2, 2, 2, 135, 371, 
	3, 2, 2, 2, 137, 373, 3, 2, 2, 2, 139, 375, 3, 2, 2, 2, 141, 377, 3, 2, 
	2, 2, 143, 379, 3, 2, 2, 2, 145, 381, 3, 2, 2, 2, 147, 383, 3, 2, 2, 2, 
	149, 385, 3, 2, 2, 2, 151, 398, 3, 2, 2, 2, 153, 445, 3, 2, 2, 2, 155, 
	447, 3, 2, 2, 2, 157, 461, 3, 2, 2, 2, 159, 160, 7, 46, 2, 2, 160, 4, 3, 
	2, 2, 2, 161, 162, 7, 66, 2, 2, 162, 163, 7, 112, 2, 2, 163, 164, 7, 99, 
	2, 2, 164, 165, 7, 111, 2, 2, 165, 166, 7, 103, 2, 2, 166, 6, 3, 2, 2, 
	2, 167, 168, 7, 66, 2, 2, 168, 169, 7, 107, 2, 2, 169, 170, 7, 102, 2, 
	2, 170, 8, 3, 2, 2, 2, 171, 172, 7, 66, 2, 2, 172, 173, 7, 102, 2, 2, 173, 
	174, 7, 103, 2, 2, 174, 175, 7, 117, 2, 2, 175, 176, 7, 101, 2, 2, 176, 
	10, 3, 2, 2, 2, 177, 178, 9, 2, 2, 2, 178, 12, 3, 2, 2, 2, 179, 180, 9, 
	3, 2, 2, 180, 14, 3, 2, 2, 2, 181, 182, 9, 4, 2, 2, 182, 16, 3, 2, 2, 2, 
	183, 184, 9, 5, 2, 2, 184, 18, 3, 2, 2, 2, 185, 186, 9, 6, 2, 2, 186, 20, 
	3, 2, 2, 2, 187, 188, 9, 7, 2, 2, 188, 22, 3, 2, 2, 2, 189, 190, 9, 8, 
	2, 2, 190, 24, 3, 2, 2, 2, 191, 192, 9, 9, 2, 2, 192, 26, 3, 2, 2, 2, 193, 
	194, 9, 10, 2, 2, 194, 28, 3, 2, 2, 2, 195, 196, 9, 11, 2, 2, 196, 30, 
	3, 2, 2, 2, 197, 198, 9, 12, 2, 2, 198, 32, 3, 2, 2, 2, 199, 200, 9, 13, 
	2, 2, 200, 34, 3, 2, 2, 2, 201, 202, 9, 14, 2, 2, 202, 36, 3, 2, 2, 2, 
	203, 204, 9, 15, 2, 2, 204, 38, 3, 2, 2, 2, 205, 206, 9, 16, 2, 2, 206, 
	40, 3, 2, 2, 2, 207, 208, 9, 17, 2, 2, 208, 42, 3, 2, 2, 2, 209, 210, 9, 
	18, 2, 2, 210, 44, 3, 2, 2, 2, 211, 212, 9, 19, 2, 2, 212, 46, 3, 2, 2, 
	2, 213, 214, 9, 20, 2, 2, 214, 48, 3, 2, 2, 2, 215, 216, 9, 21, 2, 2, 216, 
	50, 3, 2, 2, 2, 217, 218, 9, 22, 2, 2, 218, 52, 3, 2, 2, 2, 219, 220, 9, 
	23, 2, 2, 220, 54, 3, 2, 2, 2, 221, 222, 9, 24, 2, 2, 222, 56, 3, 2, 2, 
	2, 223, 224, 9, 25, 2, 2, 224, 58, 3, 2, 2, 2, 225, 226, 9, 26, 2, 2, 226, 
	60, 3, 2, 2, 2, 227, 228, 9, 27, 2, 2, 228, 62, 3, 2, 2, 2, 229, 230, 9, 
	28, 2, 2, 230, 64, 3, 2, 2, 2, 231, 233, 9, 7, 2, 2, 232, 234, 7, 47, 2, 
	2, 233, 232, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 236, 3, 2, 2, 2, 235, 
	237, 5, 11, 6, 2, 236, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 236, 
	3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 66, 3, 2, 2, 2, 240, 241, 5, 39, 
	20, 2, 241, 242, 5, 29, 15, 2, 242, 243, 5, 35, 18, 2, 243, 68, 3, 2, 2, 
	2, 244, 245, 5, 47, 24, 2, 245, 246, 5, 53, 27, 2, 246, 247, 5, 35, 18, 
	2, 247, 248, 5, 21, 11, 2, 248, 70, 3, 2, 2, 2, 249, 250, 7, 40, 2, 2, 
	250, 251, 7, 40, 2, 2, 251, 72, 3, 2, 2, 2, 252, 253, 7, 126, 2, 2, 253, 
	254, 7, 126, 2, 2, 254, 74, 3, 2, 2, 2, 255, 256, 5, 17, 9, 2, 256, 257, 
	5, 41, 21, 2, 257, 258, 5, 39, 20, 2, 258, 259, 5, 17, 9, 2, 259, 76, 3, 
	2, 2, 2, 260, 261, 5, 29, 15, 2, 261, 262, 5, 23, 12, 2, 262, 78, 3, 2, 
	2, 2, 263, 264, 5, 21, 11, 2, 264, 265, 5, 35, 18, 2, 265, 266, 5, 49, 
	25, 2, 266, 267, 5, 21, 11, 2, 267, 80, 3, 2, 2, 2, 268, 269, 5, 47, 24, 
	2, 269, 270, 5, 21, 11, 2, 270, 271, 5, 51, 26, 2, 271, 272, 5, 53, 27, 
	2, 272, 273, 5, 47, 24, 2, 273, 274, 5, 39, 20, 2, 274, 82, 3, 2, 2, 2, 
	275, 276, 5, 51, 26, 2, 276, 277, 5, 47, 24, 2, 277, 278, 5, 53, 27, 2, 
	278, 279, 5, 21, 11, 2, 279, 84, 3, 2, 2, 2, 280, 281, 5, 23, 12, 2, 281, 
	282, 5, 13, 7, 2, 282, 283, 5, 35, 18, 2, 283, 284, 5, 49, 25, 2, 284, 
	285, 5, 21, 11, 2, 285, 86, 3, 2, 2, 2, 286, 287, 5, 39, 20, 2, 287, 288, 
	5, 53, 27, 2, 288, 289, 5, 35, 18, 2, 289, 290, 5, 35, 18, 2, 290, 88, 
	3, 2, 2, 2, 291, 292, 5, 49, 25, 2, 292, 293, 5, 13, 7, 2, 293, 294, 5, 
	35, 18, 2, 294, 295, 5, 29, 15, 2, 295, 296, 5, 21, 11, 2, 296, 297, 5, 
	39, 20, 2, 297, 298, 5, 17, 9, 2, 298, 299, 5, 21, 11, 2, 299, 90, 3, 2, 
	2, 2, 300, 301, 5, 15, 8, 2, 301, 302, 5, 21, 11, 2, 302, 303, 5, 25, 13, 
	2, 303, 304, 5, 29, 15, 2, 304, 305, 5, 39, 20, 2, 305, 92, 3, 2, 2, 2, 
	306, 307, 5, 21, 11, 2, 307, 308, 5, 39, 20, 2, 308, 309, 5, 19, 10, 2, 
	309, 94, 3, 2, 2, 2, 310, 312, 9, 29, 2, 2, 311, 310, 3, 2, 2, 2, 312, 
	313, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 318, 
	3, 2, 2, 2, 315, 317, 9, 30, 2, 2, 316, 315, 3, 2, 2, 2, 317, 320, 3, 2, 
	2, 2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 96, 3, 2, 2, 2, 
	320, 318, 3, 2, 2, 2, 321, 323, 4, 50, 59, 2, 322, 321, 3, 2, 2, 2, 323, 
	324, 3, 2, 2, 2, 324, 322, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 98, 3, 
	2, 2, 2, 326, 327, 7, 45, 2, 2, 327, 100, 3, 2, 2, 2, 328, 329, 7, 47, 
	2, 2, 329, 102, 3, 2, 2, 2, 330, 331, 7, 49, 2, 2, 331, 104, 3, 2, 2, 2, 
	332, 333, 7, 44, 2, 2, 333, 106, 3, 2, 2, 2, 334, 335, 7, 63, 2, 2, 335, 
	336, 7, 63, 2, 2, 336, 108, 3, 2, 2, 2, 337, 338, 7, 64, 2, 2, 338, 110, 
	3, 2, 2, 2, 339, 340, 7, 62, 2, 2, 340, 112, 3, 2, 2, 2, 341, 342, 7, 64, 
	2, 2, 342, 343, 7, 63, 2, 2, 343, 114, 3, 2, 2, 2, 344, 345, 7, 62, 2, 
	2, 345, 346, 7, 63, 2, 2, 346, 116, 3, 2, 2, 2, 347, 348, 7, 35, 2, 2, 
	348, 349, 7, 63, 2, 2, 349, 118, 3, 2, 2, 2, 350, 351, 7, 35, 2, 2, 351, 
	120, 3, 2, 2, 2, 352, 353, 7, 60, 2, 2, 353, 354, 7, 63, 2, 2, 354, 122, 
	3, 2, 2, 2, 355, 356, 7, 63, 2, 2, 356, 124, 3, 2, 2, 2, 357, 358, 7, 45, 
	2, 2, 358, 359, 7, 63, 2, 2, 359, 126, 3, 2, 2, 2, 360, 361, 7, 47, 2, 
	2, 361, 362, 7, 63, 2, 2, 362, 128, 3, 2, 2, 2, 363, 364, 7, 44, 2, 2, 
	364, 365, 7, 63, 2, 2, 365, 130, 3, 2, 2, 2, 366, 367, 7, 49, 2, 2, 367, 
	368, 7, 63, 2, 2, 368, 132, 3, 2, 2, 2, 369, 370, 7, 93, 2, 2, 370, 134, 
	3, 2, 2, 2, 371, 372, 7, 95, 2, 2, 372, 136, 3, 2, 2, 2, 373, 374, 7, 61, 
	2, 2, 374, 138, 3, 2, 2, 2, 375, 376, 7, 125, 2, 2, 376, 140, 3, 2, 2, 
	2, 377, 378, 7, 127, 2, 2, 378, 142, 3, 2, 2, 2, 379, 380, 7, 42, 2, 2, 
	380, 144, 3, 2, 2, 2, 381, 382, 7, 43, 2, 2, 382, 146, 3, 2, 2, 2, 383, 
	384, 7, 48, 2, 2, 384, 148, 3, 2, 2, 2, 385, 393, 7, 36, 2, 2, 386, 387, 
	7, 94, 2, 2, 387, 392, 11, 2, 2, 2, 388, 389, 7, 36, 2, 2, 389, 392, 7, 
	36, 2, 2, 390, 392, 10, 31, 2, 2, 391, 386, 3, 2, 2, 2, 391, 388, 3, 2, 
	2, 2, 391, 390, 3, 2, 2, 2, 392, 395, 3, 2, 2, 2, 393, 391, 3, 2, 2, 2, 
	393, 394, 3, 2, 2, 2, 394, 396, 3, 2, 2, 2, 395, 393, 3, 2, 2, 2, 396, 
	397, 7, 36, 2, 2, 397, 150, 3, 2, 2, 2, 398, 399, 5, 95, 48, 2, 399, 400, 
	5, 147, 74, 2, 400, 401, 5, 95, 48, 2, 401, 152, 3, 2, 2, 2, 402, 404, 
	5, 11, 6, 2, 403, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 403, 3, 2, 
	2, 2, 405, 406, 3, 2, 2, 2, 406, 408, 3, 2, 2, 2, 407, 403, 3, 2, 2, 2, 
	407, 408, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 411, 7, 48, 2, 2, 410, 
	412, 5, 11, 6, 2, 411, 410, 3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 411, 
	3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 446, 3, 2, 2, 2, 415, 417, 5, 11, 
	6, 2, 416, 415, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 
	418, 419, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 421, 7, 48, 2, 2, 421, 
	422, 5, 65, 33, 2, 422, 446, 3, 2, 2, 2, 423, 425, 5, 11, 6, 2, 424, 423, 
	3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 424, 3, 2, 2, 2, 426, 427, 3, 2, 
	2, 2, 427, 429, 3, 2, 2, 2, 428, 424, 3, 2, 2, 2, 428, 429, 3, 2, 2, 2, 
	429, 430, 3, 2, 2, 2, 430, 432, 7, 48, 2, 2, 431, 433, 5, 11, 6, 2, 432, 
	431, 3, 2, 2, 2, 433, 434, 3, 2, 2, 2, 434, 432, 3, 2, 2, 2, 434, 435, 
	3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 437, 5, 65, 33, 2, 437, 446, 3, 
	2, 2, 2, 438, 440, 5, 11, 6, 2, 439, 438, 3, 2, 2, 2, 440, 441, 3, 2, 2, 
	2, 441, 439, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 
	444, 5, 65, 33, 2, 444, 446, 3, 2, 2, 2, 445, 407, 3, 2, 2, 2, 445, 416, 
	3, 2, 2, 2, 445, 428, 3, 2, 2, 2, 445, 439, 3, 2, 2, 2, 446, 154, 3, 2, 
	2, 2, 447, 448, 7, 49, 2, 2, 448, 449, 7, 49, 2, 2, 449, 453, 3, 2, 2, 
	2, 450, 452, 11, 2, 2, 2, 451, 450, 3, 2, 2, 2, 452, 455, 3, 2, 2, 2, 453, 
	454, 3, 2, 2, 2, 453, 451, 3, 2, 2, 2, 454, 456, 3, 2, 2, 2, 455, 453, 
	3, 2, 2, 2, 456, 457, 7, 12, 2, 2, 457, 458, 3, 2, 2, 2, 458, 459, 8, 78, 
	2, 2, 459, 156, 3, 2, 2, 2, 460, 462, 9, 32, 2, 2, 461, 460, 3, 2, 2, 2, 
	462, 463, 3, 2, 2, 2, 463, 461, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464, 
	465, 3, 2, 2, 2, 465, 466, 8, 79, 2, 2, 466, 158, 3, 2, 2, 2, 22, 2, 233, 
	238, 313, 316, 318, 324, 391, 393, 405, 407, 413, 418, 426, 428, 434, 441, 
	445, 453, 463, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'@name'", "'@id'", "'@desc'", "", "", "'&&'", "'||'", "", "", 
	"", "", "", "", "", "", "", "", "", "", "'+'", "'-'", "'/'", "'*'", "'=='", 
	"'>'", "'<'", "'>='", "'<='", "'!='", "'!'", "':='", "'='", "'+='", "'-='", 
	"'*='", "'/='", "'['", "']'", "';'", "'{'", "'}'", "'('", "')'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "NIL", "RULE", "AND", "OR", "CONC", "IF", "ELSE", "RETURN", 
	"TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN", "END", "SIMPLENAME", 
	"INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT", "LT", "GTE", "LTE", 
	"NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL", 
	"DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", 
	"RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", 
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "DEC_DIGIT", "A", "B", "C", "D", "E", "F", 
	"G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", 
	"V", "W", "X", "Y", "Z", "EXPONENT_NUM_PART", "NIL", "RULE", "AND", "OR", 
	"CONC", "IF", "ELSE", "RETURN", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", 
	"BEGIN", "END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", 
	"GT", "LT", "GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", 
	"MINUSEQUAL", "MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", 
	"LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", 
	"DOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
}

type gengineLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewgengineLexer(input antlr.CharStream) *gengineLexer {

	l := new(gengineLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "gengine.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gengineLexer tokens.
const (
	gengineLexerT__0 = 1
	gengineLexerT__1 = 2
	gengineLexerT__2 = 3
	gengineLexerT__3 = 4
	gengineLexerNIL = 5
	gengineLexerRULE = 6
	gengineLexerAND = 7
	gengineLexerOR = 8
	gengineLexerCONC = 9
	gengineLexerIF = 10
	gengineLexerELSE = 11
	gengineLexerRETURN = 12
	gengineLexerTRUE = 13
	gengineLexerFALSE = 14
	gengineLexerNULL_LITERAL = 15
	gengineLexerSALIENCE = 16
	gengineLexerBEGIN = 17
	gengineLexerEND = 18
	gengineLexerSIMPLENAME = 19
	gengineLexerINT = 20
	gengineLexerPLUS = 21
	gengineLexerMINUS = 22
	gengineLexerDIV = 23
	gengineLexerMUL = 24
	gengineLexerEQUALS = 25
	gengineLexerGT = 26
	gengineLexerLT = 27
	gengineLexerGTE = 28
	gengineLexerLTE = 29
	gengineLexerNOTEQUALS = 30
	gengineLexerNOT = 31
	gengineLexerASSIGN = 32
	gengineLexerSET = 33
	gengineLexerPLUSEQUAL = 34
	gengineLexerMINUSEQUAL = 35
	gengineLexerMULTIEQUAL = 36
	gengineLexerDIVEQUAL = 37
	gengineLexerLSQARE = 38
	gengineLexerRSQARE = 39
	gengineLexerSEMICOLON = 40
	gengineLexerLR_BRACE = 41
	gengineLexerRR_BRACE = 42
	gengineLexerLR_BRACKET = 43
	gengineLexerRR_BRACKET = 44
	gengineLexerDOT = 45
	gengineLexerDQUOTA_STRING = 46
	gengineLexerDOTTEDNAME = 47
	gengineLexerREAL_LITERAL = 48
	gengineLexerSL_COMMENT = 49
	gengineLexerWS = 50
)

