package wadfile

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"ltk/helper"
	"testing"
	"time"
)

func init() {

}

/*

public string GetInfo()
        {
            return this.Path + '\n'
                + "Compression Type: " + this.Entry.Type.ToString() + '\n'
                + "Compressed Size: " + this.Entry.CompressedSize + '\n'
                + "Uncompressed Size: " + this.Entry.UncompressedSize + '\n'
                + "Checksum: " + UtilitiesFantome.ByteArrayToHex(this.Entry.Checksum);
        }

public static string ByteArrayToHex(byte[] array)
{
char[] chArray = new char[array.Length * 2];
for (int index = 0; index < array.Length; ++index)
{
int num1 = (int) array[index] >> 4;
chArray[index * 2] = (char) (55 + num1 + (num1 - 10 >> 31 & -7));
int num2 = (int) array[index] & 15;
chArray[index * 2 + 1] = (char) (55 + num2 + (num2 - 10 >> 31 & -7));
}
return new string(chArray);
}*/

// []byte to xxhash
// entry.XXHash.ToString("x16")  251074681296703432 => 037bff17a6c8cfc8
// return string.Format("{0}.{1}", entry.XXHash.ToString("x16"), extension);
func Test_ByteArrayToHex(t *testing.T) {
	var array = []byte{102, 148, 61, 186, 77, 234, 68, 134} // 8 bytes

	// char[] chArray = new char[array.Length * 2];
	ch := make([]byte, len(array)*2)
	t.Log(ch)

	for i, b := range array {
		// int num1 = (int) array[index] >> 4;
		num1 := b >> 4
		// chArray[index * 2] = (char) (55 + num1 + (num1 - 10 >> 31 & -7));
		ch[i*2] = byte(55 + num1 + (num1 - 10>>31&-7))
		// int num2 = (int) array[index] & 15;
		num2 := b & 15
		// chArray[index * 2 + 1] = (char) (55 + num2 + (num2 - 10 >> 31 & -7));
		ch[i*2+1] = byte(55 + num2 + (num2 - 10>>31&-7))
	}
	// new string(chArray);
	t.Log(string(ch))

	var res = []byte{34, 115, 103, 77, 56, 116, 78, 97, 102, 71, 57, 52, 61, 34} // 16 bytes
	t.Log(res)

	// 这样转直接能对得上
	var str = fmt.Sprintf("%x", array)
	t.Log(str)

	toString := hex.EncodeToString(array)
	t.Log(toString)

	chars := []byte(str)

	var v byte
	v = 'a'
	t.Log(v)

	// "sgM8tNafG94="
	t.Log(chars)

}

func TestRead(t *testing.T) {
	filepath := "../../files/wad/Aatrox.wad.client"

	start := time.Now().Unix()
	wad, err := Read(filepath)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Read time:", time.Now().Unix()-start)

	fmt.Println(wad.Signature())

	defer wad.File.Close()
	s, _ := json.Marshal(wad.signature)
	fmt.Println(string(s))

	for xxhash, entry := range wad.Entries {
		bytes, err := NewWadEntryDataHandle(entry).GetDecompressedBytes()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(helper.HashToHex16(xxhash), len(bytes))
	}
	fmt.Println(wad.FileCount)
}
