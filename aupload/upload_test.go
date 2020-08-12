package aupload

import (
    "fmt"
    "github.com/asktop/gotools/arand"
    "github.com/asktop/gotools/astring"
    "testing"
)

func TestUploadFromPath(t *testing.T) {
    client, err := NewClient(nil, DriverLocal, Config{Local: LocalConfig{Site: "http://127.0.0.1:8881", Bucket: "upload"}})
    if err != nil {
        fmt.Println(err)
        return
    }

    fInfo, err := client.UploadFromPath(`D:\00Down\cache\1.jpg`, astring.JoinURL("test", arand.RandMd5()))
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(fInfo.OldName)
        fmt.Println(fInfo.Url)
        fmt.Println(fInfo.Path)
    }
}

func TestDeleteFile(t *testing.T) {
    client, err := NewClient(nil, DriverLocal, Config{Local: LocalConfig{Site: "http://127.0.0.1:8881", Bucket: "upload"}})
    if err != nil {
        fmt.Println(err)
        return
    }

    err = client.DeleteFile(`https://asktop-1258252583.cos.ap-guangzhou.myqcloud.com/test/f295ee56a7ce4c4ad112501f14035390.jpg`)
    if err != nil {
        fmt.Println(err)
    }
}
