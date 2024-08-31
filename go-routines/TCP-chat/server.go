package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newSever() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command), // go channel
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_WILL:
			s.will(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.lsrooms(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		default:
		}
	}
}

func (s *server) newClient(conn *net.Conn) {
	log.Printf("New cliented connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		will:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) will(c *client, args []string) {
	c.will = args[1]
	c.msg(fmt.Sprintf("Your username is: %s", c.will))
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)

	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the chat", c.will))
	c.msg(fmt.Sprintf("Welcome to %s", r.name))
}

func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("Kindly join a room first"))
		return
	}
	c.room.broadcast(c, c.will + ": " + strings.Join(args[1:len(args)], " "))
}

func (s *server) lsrooms(c *client, args []string) {
	var rooms []string

	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("Availbale rooms are: %s", strings.Join(rooms, ", ")))
}

func (s *server) quit(c *client, args []string) {
log.Fatalf("%s has disconnected:", c.conn.RemoteAddr().String())
s.quitCurrentRoom(c)
c.msg("See you soon. ;)")
c.conn.Close()

}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the chat", c.will))
	}
}
