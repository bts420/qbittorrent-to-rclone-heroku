package main


import (
        "fmt"
        "sync"
        "os/exec"
        "os"
        "net/http"
)


//
//FUNC QB
//

func qb() {
// raw file

r, e := http.Get("https://raw.githubusercontent.com/bts420/qbittorrent-to-rclone-heroku/main/qb-fun-qb-run-qb-only")
   if e != nil {
      panic(e)
   }
   defer r.Body.Close()
   f, e := os.Create("./scriptplusconf/qb.sh")
   if e != nil {
      panic(e)
   }
   defer f.Close()
   f.ReadFrom(r.Body)

err := os.Chmod("./scriptplusconf/qb.sh", 0777)
        if err != nil {
                fmt.Println(err)
        }
// EXECUTE
        cmd := exec.Command("bash", "./scriptplusconf/qb.sh")
        cmd.Run()
}




//
//FUNC QB
//

func chqb() {
// raw file

r, e := http.Get("https://raw.githubusercontent.com/bts420/qbittorrent-to-rclone-heroku/main/qb-fun-qb-change-qb-only")
   if e != nil {
      panic(e)
   }
   defer r.Body.Close()
   f, e := os.Create("./scriptplusconf/chqb.sh")
   if e != nil {
      panic(e)
   }
   defer f.Close()
   f.ReadFrom(r.Body)

err := os.Chmod("./scriptplusconf/chqb.sh", 0777)
        if err != nil {
                fmt.Println(err)
        }
// EXECUTE
        cmd := exec.Command("bash", "./scriptplusconf/chqb.sh")
        cmd.Run()
}





//
//FUNC QB
//


func ng() {
// raw file

r, e := http.Get("https://raw.githubusercontent.com/bts420/qbittorrent-to-rclone-heroku/main/qb-heroku-fun-ng")
   if e != nil {
      panic(e)
   }
   defer r.Body.Close()
   f, e := os.Create("./scriptplusconf/ng.sh")
   if e != nil {
      panic(e)
   }
   defer f.Close()
   f.ReadFrom(r.Body)

err := os.Chmod("./scriptplusconf/ng.sh", 0777)
        if err != nil {
                fmt.Println(err)
        }
// EXECUTE
        cmd := exec.Command("bash", "./scriptplusconf/ng.sh")
        cmd.Run()
}


//
//FUNC QB
//

func rc() {
// raw file

r, e := http.Get("https://raw.githubusercontent.com/bts420/qbittorrent-to-rclone-heroku/main/scriptplusconf/addon.sh")
   if e != nil {
      panic(e)
   }
   defer r.Body.Close()
   f, e := os.Create("./scriptplusconf/rc.sh")
   if e != nil {
      panic(e)
   }
   defer f.Close()
   f.ReadFrom(r.Body)

err := os.Chmod("./scriptplusconf/rc.sh", 0777)
        if err != nil {
                fmt.Println(err)
        }
// EXECUTE
        cmd := exec.Command("bash", "./scriptplusconf/rc.sh")
        cmd.Run()
}



func main() {
        os.MkdirAll("./scriptplusconf", os.ModePerm)

        var wg sync.WaitGroup


        wg.Add( 4 )
        go qb()
        go ng()
        go rc()
	go chqb()
        wg.Wait()
        fmt.Printf( "All process Dead, NOT HEROKU\n" )

}
