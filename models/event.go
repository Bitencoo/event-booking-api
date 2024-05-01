package models

import (
	"time"

	"github.com/Bitencoo/event-booking-api.git/db"
)

type Event struct {
	ID int64
	Name string `binding: required`
	Description string `binding: required`
	Location string `binding: required`
	DateTime time.Time `binding: required`
	UserID int64
}

func (e *Event) Save() error {
	query := `
		INSERT INTO EVENTS(name, description, location, datetime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err;
	}
	
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err;
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err;
	}

	(*e).ID = id
	
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM EVENTS"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err;
	}

	defer rows.Close()

	var events = []Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
		SELECT * FROM EVENTS WHERE ID = ?
	`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) UpdateEvent() error{
	query := `
		UPDATE EVENTS 
		SET name = ?, description = ?, location = ?, datetime = ?
		WHERE id = ? 
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func DeleteEvent(id int64) (error) {
	query := `
		DELETE FROM EVENTS WHERE ID = ?
	`

	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

func (e Event) Register(userId int64) (int64, error) {
	query := `INSERT INTO REGISTRATIONS(event_id, user_id) VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.ID, userId)

	if err != nil {
		return 0, err
	}

	registrationID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return registrationID, nil
}

func (e Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM REGISTRATIONS WHERE event_id = ? AND user_id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)

	return err
}

