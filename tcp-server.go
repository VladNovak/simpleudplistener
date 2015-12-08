package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing
import "os"

func main() {

  args := os.Args  

  if (len(args) > 1) && (args[1] == "udp") {
     fmt.Println("Launching server in udp mode...")
     ln, _ := net.ResolveUDPAddr("udp", "8010")
    
     /* Now listen at selected port */
     conn, _ := net.ListenUDP("udp", ln)
     defer conn.Close()
      
     buf := make([]byte, 1024)

     for {
        n,addr,_ := conn.ReadFromUDP(buf)
        message := string(buf[0:n])
        newmessage := strings.ToUpper(message[:len(message)-1])
        if newmessage == "STATUS" {
          conn.WriteTo([]byte("OK\n"),addr)
        } else if newmessage == "TERMINATE" {
          break
        } else {
          // send new string back to client
          conn.WriteTo([]byte(newmessage + "\n"),addr)
        }
     }

  } else  {
      fmt.Println("Launching server...")
     ln, _ := net.Listen("tcp", "8010")

     // accept connection on port
     for {
      conn, _ := ln.Accept()
      for {
       // will listen for message to process ending in newline (\n)
       message, _ := bufio.NewReader(conn).ReadString('\n')
       newmessage := strings.ToUpper(message[:len(message)-1])
       if newmessage == "STATUS" {
          conn.Write([]byte("OK\n"))
       } else if newmessage == "TERMINATE" {
          break
       } else {
          // send new string back to client
          conn.Write([]byte(newmessage + "\n"))
       }
     }
     } 
  } 
}
