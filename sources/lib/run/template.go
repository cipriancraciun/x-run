
package zrun


import "encoding/base64"
import "encoding/hex"
import "encoding/json"
import "fmt"
import "io"
import "os"
import "path"
import "path/filepath"
import "strconv"
import "strings"
import "text/template"

import . "github.com/volution/z-run/lib/library"
import . "github.com/volution/z-run/lib/mainlib"
import . "github.com/volution/z-run/lib/common"




func PrintMain (_selfExecutable string, _arguments []string, _environment map[string]string) (*Error) {
	
	if len (_arguments) != 1 {
		return Errorf (0xbc2c2406, "invalid arguments")
	}
	
	var _sourcePath = _arguments[0]
	
	var _sourceBody string
	if _, _data, _error := loadFromFile (_sourcePath); _error == nil {
		_sourceBody = string (_data)
	} else {
		return _error
	}
	
	_error := executePrint_0 (_sourceBody, os.Stdout)
	if _error != nil {
		return _error
	}
	
	panic (ExitMainSucceeded ())
}


func executePrint (_library LibraryStore, _scriptlet *Scriptlet, _context *Context, _output io.Writer) (*Error) {
	
	if _scriptlet.Interpreter != "<print>" {
		return Errorf (0x43c4a524, "invalid interpreter")
	}
	
	if len (_context.cleanArguments) != 0 {
		return Errorf (0x4f9e0e3b, "invalid arguments")
	}
	
	return executePrint_0 (_scriptlet.Body, _output)
}


func executePrint_0 (_source string, _output io.Writer) (*Error) {
	
	if _, _error := _output.Write ([]byte (_source)); _error != nil {
		return Errorw (0x1fea6849, _error)
	}
	
	return nil
}




func TemplateMain (_selfExecutable string, _arguments []string, _environment map[string]string) (*Error) {
	
	if len (_arguments) < 1 {
		return Errorf (0x47c3f9f1, "invalid arguments")
	}
	
	var _sourcePath = _arguments[0]
	_arguments = _arguments[1:]
	
	var _sourceBody string
	if _, _data, _error := loadFromFile (_sourcePath); _error == nil {
		_sourceBody = string (_data)
	} else {
		return _error
	}
	
	_error := executeTemplate_0 (
			_sourceBody,
			_arguments,
			_environment,
			_selfExecutable,
			"",
			"",
			"",
			"",
			nil,
			os.Stdout,
		)
	if _error != nil {
		return _error
	}
	
	panic (ExitMainSucceeded ())
}




func executeTemplate (_library LibraryStore, _scriptlet *Scriptlet, _context *Context, _output io.Writer) (*Error) {
	
	if _scriptlet.Interpreter != "<template>" {
		return Errorf (0xa18a5ca9, "invalid interpreter")
	}
	
	_libraryUrl := _library.Url ()
	_libraryIdentifier := ""
	if _libraryIdentifier_0, _error := _library.Identifier (); _error == nil {
		_libraryIdentifier = _libraryIdentifier_0
	} else {
		return _error
	}
	_libraryFingerprint := ""
	if _libraryFingerprint_0, _error := _library.Fingerprint (); _error == nil {
		_libraryFingerprint = _libraryFingerprint_0
	} else {
		return _error
	}
	
	_extraFunctions := make (map[string]interface{}, 16)
	
	_extraFunctions["Z_zspawn_capture"] = func (_scriptlet string, _arguments ... string) (string, error) {
			return templateFuncZrun (_library, _context, _scriptlet, _arguments)
		}
	
	_extraFunctions["Z_select_top_labels"] = func () ([]string, error) {
			if _labels, _error := _library.SelectLabels (); _error == nil {
				return _labels, nil
			} else {
				return nil, _error.ToError ()
			}
		}
	_extraFunctions["Z_select_all_labels"] = func () ([]string, error) {
			if _labels, _error := _library.SelectLabelsAll (); _error == nil {
				return _labels, nil
			} else {
				return nil, _error.ToError ()
			}
		}
	_extraFunctions["Z_select_all_fingerprints"] = func () ([]string, error) {
			if _fingerprints, _error := _library.SelectFingerprints (); _error == nil {
				return _fingerprints, nil
			} else {
				return nil, _error.ToError ()
			}
		}
	
	_extraFunctions["Z_resolve_full_by_label"] = func (_scriptlet string) (*Scriptlet, error) {
			if _scriptlet, _error := _library.ResolveFullByLabel (_scriptlet); _error == nil {
				return _scriptlet, nil
			} else {
				return nil, _error.ToError ()
			}
		}
	_extraFunctions["Z_resolve_full_by_label"] = func (_scriptlet string) (*Scriptlet, error) {
			if _scriptlet, _error := _library.ResolveFullByLabel (_scriptlet); _error == nil {
				return _scriptlet, nil
			} else {
				return nil, _error.ToError ()
			}
		}
	_extraFunctions["Z_resolve_full_by_fingerprint"] = func (_scriptlet string) (*Scriptlet, error) {
			if _scriptlet, _error := _library.ResolveFullByFingerprint (_scriptlet); _error == nil {
				return _scriptlet, nil
			} else {
				return nil, _error.ToError ()
			}
		}
	
	_extraFunctions["Z_resolve_meta_by_label"] = func (_scriptlet string) (*Scriptlet, error) {
			if _scriptlet, _error := _library.ResolveMetaByLabel (_scriptlet); _error == nil {
				return _scriptlet, nil
			} else {
				return nil, _error.ToError ()
			}
		}
	_extraFunctions["Z_resolve_meta_by_fingerprint"] = func (_scriptlet string) (*Scriptlet, error) {
			if _scriptlet, _error := _library.ResolveMetaByFingerprint (_scriptlet); _error == nil {
				return _scriptlet, nil
			} else {
				return nil, _error.ToError ()
			}
		}
	
	_extraFunctions["Z_resolve_body_by_label"] = func (_scriptlet string) (string, error) {
			if _body, _, _error := _library.ResolveBodyByLabel (_scriptlet); _error == nil {
				return _body, nil
			} else {
				return "", _error.ToError ()
			}
		}
	_extraFunctions["Z_resolve_body_by_fingerprint"] = func (_scriptlet string) (string, error) {
			if _body, _, _error := _library.ResolveBodyByFingerprint (_scriptlet); _error == nil {
				return _body, nil
			} else {
				return "", _error.ToError ()
			}
		}
	_extraFunctions["Z_resolve_fingerprint_by_label"] = func (_scriptlet string) (string, error) {
			if _fingerprint, _, _error := _library.ResolveFingerprintByLabel (_scriptlet); _error == nil {
				return _fingerprint, nil
			} else {
				return "", _error.ToError ()
			}
		}
	
	return executeTemplate_0 (
			_scriptlet.Body,
			_context.cleanArguments,
			_context.cleanEnvironment,
			_context.selfExecutable,
			_context.workspace,
			_libraryUrl,
			_libraryIdentifier,
			_libraryFingerprint,
			_extraFunctions,
			_output,
		)
}




func executeTemplate_0 (
			_source string,
			_arguments []string,
			_environment map[string]string,
			_selfExecutable string,
			_workspace string,
			_libraryUrl string,
			_libraryIdentifier string,
			_libraryFingerprint string,
			_extraFunctions map[string]interface{},
			_output io.Writer,
		) (*Error) {
	
	_functions := templateFunctions ()
	for _name, _function := range _extraFunctions {
		_functions[_name] = _function
	}
	
	_template := template.New ("z-run")
	_template.Funcs (_functions)
	if _, _error := _template.Parse (_source); _error != nil {
		return Errorw (0xad3804cc, _error)
	}
	
	_workspaceIdentifier := NewFingerprinter () .String (_workspace) .Build ()
	
	_data := map[string]interface{} {
			"arguments" : _arguments,
			"environment" : _environment,
			"ZRUN_EXECUTABLE" : _selfExecutable,
			"ZRUN_WORKSPACE" : _workspace,
			"ZRUN_WORKSPACE_IDENTIFIER" : _workspaceIdentifier,
			"ZRUN_LIBRARY_URL" : _libraryUrl,
			"ZRUN_LIBRARY_IDENTIFIER" : _libraryIdentifier,
			"ZRUN_LIBRARY_FINGERPRINT" : _libraryFingerprint,
		}
	
	if _error := _template.Execute (_output, _data); _error != nil {
		return Errorw (0x0d6d4b96, _error)
	}
	
	return nil
}




func templateFuncZrun (_library LibraryStore, _context *Context, _scriptletLabel string, _arguments []string) (string, error) {
	if strings.HasPrefix (_scriptletLabel, ":: ") {
		_scriptletLabel = _scriptletLabel[3:]
	}
	_libraryIdentifier := ""
	if _libraryIdentifier_0, _error := _library.Identifier (); _error == nil {
		_libraryIdentifier = _libraryIdentifier_0
	} else {
		return "", _error.ToError ()
	}
	_libraryFingerprint := ""
	if _libraryFingerprint_0, _error := _library.Fingerprint (); _error == nil {
		_libraryFingerprint = _libraryFingerprint_0
	} else {
		return "", _error.ToError ()
	}
	if _scriptlet, _error := _library.ResolveFullByLabel (_scriptletLabel); _error == nil {
		if _scriptlet != nil {
			if _, _output, _error := loadFromScriptlet (_library.Url (), _libraryIdentifier, _libraryFingerprint, "", _scriptlet, _context); _error == nil {
				return string (_output), nil
			} else {
				return "", _error.ToError ()
			}
		} else {
			return "", Errorf (0x944c3172, "unknown scriptlet `%s`", _scriptletLabel) .ToError ()
		}
	} else {
		return "", _error.ToError ()
	}
}




func templateFunctions () (map[string]interface{}) {
	return map[string]interface{} {
			
			
			// --------------------------------------------------------------------------------
			
			
			"json_encode" : func (_input interface{}) (string, error) {
					_output, _error := json.Marshal (_input)
					return string (_output), _error
				},
			"json_decode" : func (_input string) (interface{}, error) {
					var _output interface{}
					_error := json.Unmarshal ([]byte (_input), &_output)
					return _output, _error
				},
			
			"hex_encode" : func (_input string) (string) {
					return hex.EncodeToString ([]byte (_input))
				},
			"hex_decode" : func (_input string) (string, error) {
					_output, _error := hex.DecodeString (_input)
					return string (_output), _error
				},
			
			"base64_encode" : func (_input string) (string) {
					return base64.StdEncoding.EncodeToString ([]byte (_input))
				},
			"base64_decode" : func (_input string) (string, error) {
					_output, _error := base64.StdEncoding.DecodeString (_input)
					return string (_output), _error
				},
			
			
			// --------------------------------------------------------------------------------
			
			
			"atoi" : func (_input string) (int, error) {
					_output, _error := strconv.Atoi (_input)
					return _output, _error
				},
			"itoa" : func (_input int) (string) {
					return strconv.Itoa (_input)
				},
			
			"int_parse" : func (_base int, _input string) (int, error) {
					_output, _error := strconv.ParseInt (_input, _base, 32)
					return int (_output), _error
				},
			"int_format" : func (_base int, _input int) (string) {
					return strconv.FormatInt (int64 (_input), _base)
				},
			"uint_parse" : func (_base int, _input string) (uint, error) {
					_output, _error := strconv.ParseUint (_input, _base, 32)
					return uint (_output), _error
				},
			"uint_format" : func (_base int, _input uint) (string) {
					return strconv.FormatUint (uint64 (_input), _base)
				},
			
			"int8_parse" : func (_base int, _input string) (int8, error) {
					_output, _error := strconv.ParseInt (_input, _base, 8)
					return int8 (_output), _error
				},
			"int8_format" : func (_base int, _input int8) (string) {
					return strconv.FormatInt (int64 (_input), _base)
				},
			"uint8_parse" : func (_base int, _input string) (uint8, error) {
					_output, _error := strconv.ParseUint (_input, _base, 8)
					return uint8 (_output), _error
				},
			"uint8_format" : func (_base int, _input uint8) (string) {
					return strconv.FormatUint (uint64 (_input), _base)
				},
			
			"int16_parse" : func (_base int, _input string) (int16, error) {
					_output, _error := strconv.ParseInt (_input, _base, 16)
					return int16 (_output), _error
				},
			"int16_format" : func (_base int, _input int16) (string) {
					return strconv.FormatInt (int64 (_input), _base)
				},
			"uint16_parse" : func (_base int, _input string) (uint16, error) {
					_output, _error := strconv.ParseUint (_input, _base, 16)
					return uint16 (_output), _error
				},
			"uint16_format" : func (_base int, _input uint16) (string) {
					return strconv.FormatUint (uint64 (_input), _base)
				},
			
			"int32_parse" : func (_base int, _input string) (int32, error) {
					_output, _error := strconv.ParseInt (_input, _base, 32)
					return int32 (_output), _error
				},
			"int32_format" : func (_base int, _input int32) (string) {
					return strconv.FormatInt (int64 (_input), _base)
				},
			"uint32_parse" : func (_base int, _input string) (uint32, error) {
					_output, _error := strconv.ParseUint (_input, _base, 32)
					return uint32 (_output), _error
				},
			"uint32_format" : func (_base int, _input uint32) (string) {
					return strconv.FormatUint (uint64 (_input), _base)
				},
			
			"int64_parse" : func (_base int, _input string) (int64, error) {
					_output, _error := strconv.ParseInt (_input, _base, 64)
					return int64 (_output), _error
				},
			"int64_format" : func (_base int, _input int64) (string) {
					return strconv.FormatInt (int64 (_input), _base)
				},
			"uint64_parse" : func (_base int, _input string) (uint64, error) {
					_output, _error := strconv.ParseUint (_input, _base, 64)
					return uint64 (_output), _error
				},
			"uint64_format" : func (_base int, _input uint64) (string) {
					return strconv.FormatUint (uint64 (_input), _base)
				},
			
			"float32_parse" : func (_input string) (float32, error) {
					_output, _error := strconv.ParseFloat (_input, 32)
					return float32 (_output), _error
				},
			"float32_format" : func (_input float32) (string) {
					return strconv.FormatFloat (float64 (_input), 'f', -1, 32)
				},
			
			"float64_parse" : func (_input string) (float64, error) {
					_output, _error := strconv.ParseFloat (_input, 64)
					return _output, _error
				},
			"float64_format" : func (_input float64) (string) {
					return strconv.FormatFloat (_input, 'f', -1, 64)
				},
			
			"bool_parse" : func (_input string) (bool, error) {
					_output, _error := strconv.ParseBool (_input)
					return _output, _error
				},
			"bool_format" : func (_input bool) (string) {
					return strconv.FormatBool (_input)
				},
			
			
			// --------------------------------------------------------------------------------
			
			
			"format" : func (_format string, _arguments ... interface{}) (string) {
					return fmt.Sprintf (_format, _arguments...)
				},
			
			
			// --------------------------------------------------------------------------------
			
			
			"array" : func (_inputs ... interface{}) ([]interface{}) {
					return _inputs
				},
			
			"array_append" : func (_array []interface{}, _inputs ... interface{}) ([]interface{}) {
					return append (_array, _inputs...)
				},
			
			"array_join" : func (_separator string, _input_0 []interface{}) (string, error) {
					_input := make ([]string, len (_input_0))
					for _index, _input_0 := range _input_0 {
						if _input_0, _ok := _input_0.(string); _ok {
							_input[_index] = _input_0
						} else {
							return "", Errorf (0xa2880bb1, "invalid value") .ToError ()
						}
					}
					return strings.Join (_input, _separator), nil
				},
			
			"array_contains" : func (_object interface{}, _inputs []interface{}) (bool) {
					for _, _input := range _inputs {
						if _object == _input {
							return true
						}
					}
					return false
				},
			
			
			// --------------------------------------------------------------------------------
			
			
			"join" : func (_separator string, _input []string) (string) {
					return strings.Join (_input, _separator)
				},
			
			"split_all" : func (_separator string, _input string) ([]string) {
					return strings.Split (_input, _separator)
				},
			"split_first" : func (_separator string, _input string) ([]string) {
					return strings.SplitN (_input, _separator, 2)
				},
			"split_first_n" : func (_separator string, _input string, _count int) ([]string) {
					return strings.SplitN (_input, _separator, _count)
				},
			
			"replace_all" : func (_search string, _replacement string, _input string) (string) {
					return strings.ReplaceAll (_input, _search, _replacement)
				},
			"replace_first" : func (_search string, _replacement string, _input string) (string) {
					return strings.Replace (_input, _search, _replacement, 1)
				},
			"replace_first_n" : func (_search string, _replacement string, _input string, _count int) (string) {
					return strings.Replace (_input, _search, _replacement, _count)
				},
			
			"contains" : func (_string string, _input string) (bool) {
					return strings.Contains (_input, _string)
				},
			"contains_any" : func (_characters string, _input string) (bool) {
					return strings.ContainsAny (_input, _characters)
				},
			"has_prefix" : func (_prefix string, _input string) (bool) {
					return strings.HasPrefix (_input, _prefix)
				},
			"has_suffix" : func (_suffix string, _input string) (bool) {
					return strings.HasSuffix (_input, _suffix)
				},
			"first_index_of" : func (_string string, _input string) (int) {
					return strings.Index (_input, _string)
				},
			"first_index_of_any" : func (_characters string, _input string) (int) {
					return strings.IndexAny (_input, _characters)
				},
			"last_index_of" : func (_string string, _input string) (int) {
					return strings.LastIndex (_input, _string)
				},
			"last_index_of_any" : func (_characters string, _input string) (int) {
					return strings.LastIndex (_input, _characters)
				},
			"count_of" : func (_string string, _input string) (int) {
					return strings.Count (_input, _string)
				},
			
			"repeat" : func (_count int, _input string) (string) {
					return strings.Repeat (_input, _count)
				},
			"trim_prefix" : func (_prefix string, _input string) (string) {
					return strings.TrimPrefix (_input, _prefix)
				},
			"trim_suffix" : func (_suffix string, _input string) (string) {
					return strings.TrimSuffix (_input, _suffix)
				},
			"trim_space" : func (_input string) (string) {
					return strings.TrimSpace (_input)
				},
			"trim_any" : func (_characters string, _input string) (string) {
					return strings.Trim (_input, _characters)
				},
			"trim_prefix_any" : func (_characters string, _input string) (string) {
					return strings.TrimLeft (_input, _characters)
				},
			"trim_suffix_any" : func (_characters string, _input string) (string) {
					return strings.TrimRight (_input, _characters)
				},
			
			"to_lower_ascii" : func (_input string) (string) {
					return strings.ToLower (_input)
				},
			"to_upper_ascii" : func (_input string) (string) {
					return strings.ToUpper (_input)
				},
			"to_lower" : func (_input string) (string) {
					return strings.ToLowerSpecial (nil, _input)
				},
			"to_upper" : func (_input string) (string) {
					return strings.ToUpperSpecial (nil, _input)
				},
			"to_utf8" : func (_input string) (string) {
					return strings.ToValidUTF8 (_input, "\ufffd")
				},
			
			
			// --------------------------------------------------------------------------------
			
			
			"path_join" : func (_paths ... string) (string) {
					return path.Join (_paths ...)
				},
			
			"path_dirname" : func (_path string) (string) {
					return path.Dir (_path)
				},
			"path_basename" : func (_path string) (string) {
					return path.Base (_path)
				},
			"path_split_last" : func (_path string) ([2]string) {
					_dirname, _basename := filepath.Split (_path)
					return [2]string { _dirname, _basename }
				},
			"path_split_all" : func (_path string) ([]string) {
					return filepath.SplitList (_path)
				},
			
			"path_extension" : func (_path string) (string) {
					return path.Ext (_path)
				},
			"path_without_extension" : func (_path string) (string) {
					_extension := path.Ext (_path)
					if _extension == "" {
						return _path
					}
					return _path[: len (_path) - len (_extension)]
				},
			
			"path_normalize" : func (_path string) (string) {
					return path.Clean (_path)
				},
			"path_absolute" : func (_path string) (string, error) {
					return filepath.Abs (_path)
				},
			"path_canonical" : func (_path string) (string, error) {
					return filepath.EvalSymlinks (_path)
				},
			"path_relative" : func (_base string, _path string) (string, error) {
					return filepath.Abs (_path)
				},
			
			"path_matches" : func (_pattern string, _path string) (bool, error) {
					return path.Match (_pattern, _path)
				},
			
			
			// --------------------------------------------------------------------------------
			
			
			"split_lines" : func (_input string) ([]string, error) {
					if _input == "" {
						return []string {}, nil
					}
					_array := make ([]string, 0, 128)
					_wasEmpty := false
					for _, _line := range strings.Split (_input, "\n") {
						if len (_line) > 0 {
							_array = append (_array, _line)
							_wasEmpty = false
						} else {
							_wasEmpty = true
						}
					}
					if !_wasEmpty {
						return nil, Errorf (0x1e677d43, "expected `\\n` at end of input") .ToError ()
					}
					return _array, nil
				},
			
			"join_lines" : func (_input []string) (string) {
					if len (_input) > 0 {
						return strings.Join (_input, "\n") + "\n"
					} else {
						return ""
					}
				},
			
			
			// --------------------------------------------------------------------------------
			
			
			"split_nulls" : func (_input string) ([]string, error) {
					if _input == "" {
						return []string {}, nil
					}
					_array := make ([]string, 0, 128)
					_wasEmpty := false
					for _, _line := range strings.Split (_input, "\x00") {
						if len (_line) > 0 {
							_array = append (_array, _line)
							_wasEmpty = false
						} else {
							_wasEmpty = true
						}
					}
					if !_wasEmpty {
						return nil, Errorf (0x88a1e3db, "expected `\\0` at end of input") .ToError ()
					}
					return _array, nil
				},
			
			"join_nulls" : func (_input []string) (string) {
					if len (_input) > 0 {
						return strings.Join (_input, "\x00") + "\x00"
					} else {
						return ""
					}
				},
			
			
			// --------------------------------------------------------------------------------
			
			
			"shell_quote" : func (_input string) (string) {
					// NOTE:  https://github.com/python/cpython/blob/3.8/Lib/shlex.py#L330
					return `'` + strings.ReplaceAll (_input, `'`, `'\''`) + `'`
				},
			
			"python_quote" : func (_input string) (string) {
					// NOTE:  https://docs.python.org/3/reference/lexical_analysis.html#string-and-bytes-literals
					return strconv.QuoteToASCII (_input)
				},
			
			"go_quote" : func (_input string) (string) {
					// NOTE:  https://golang.org/ref/spec#String_literals
					return strconv.QuoteToASCII (_input)
				},
			
			
			// --------------------------------------------------------------------------------
		}
}

