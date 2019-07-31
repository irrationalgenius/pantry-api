package repository

import (
	"database/sql"
	"fmt"
	"log"
	"pantry2/models"
)

//GuestRepository : GuestRepository
type GuestRepository struct{}

//GetGuests : GetGuests
func (g GuestRepository) GetGuests(db *sql.DB, guest models.Guest, guests []models.Guest) ([]models.Guest, error) {
	log.Println("Accessing the Get Guests Repository")

	rows, err := db.Query(`SELECT id, date_enrolled, status, first_name, last_name, gender,
																unit_num, st_address, state, city, zip, tel_num, email,
																count_children, count_adults, worship_place, is_member, is_baptized,
																is_espanol, is_unemployed, is_homeless, is_family,
																is_contact_ok, allergies, notes, last_date_updated
												 FROM pantry.guests`)

	if err != nil {
		return []models.Guest{}, err
	}

	for rows.Next() {
		err = rows.Scan(&guest.ID, &guest.DateEnrolled, &guest.Status, &guest.FirstName, &guest.LastName, &guest.Gender,
			&guest.UnitNum, &guest.StAddress, &guest.State, &guest.City, &guest.Zip, &guest.TelNum, &guest.Email,
			&guest.ChildNum, &guest.AdultNum, &guest.PlaceOfWorship, &guest.IsMember, &guest.IsBaptized,
			&guest.IsEspanol, &guest.IsUnemployed, &guest.IsHomeless, &guest.IsFamily,
			&guest.IsContactOk, &guest.Allergies, &guest.Notes, &guest.LastDateUpdated)

		guests = append(guests, guest)
	}

	if err != nil {
		return []models.Guest{}, err
	}

	return guests, nil
}

//AddGuest : AddGuest
func (g GuestRepository) AddGuest(db *sql.DB, guest models.Guest) error {
	log.Println("Accessing the Add Guest Repository")

	result, err := db.Exec(`INSERT INTO pantry.guests(
																	status, first_name, last_name, gender,
																	unit_num, st_address, state, city, zip, tel_num, email,
																	count_children, count_adults, worship_place, is_member, is_baptized,
																	is_espanol, is_unemployed, is_homeless, is_family,
																	is_contact_ok, allergies, notes)
                          VALUES($2, $3, $4, $5,
																 $6, $7, $8, $9, $10, $11, $12,
																 $13, $14, $15, $16, $17,
																 $18, $19, $20, $21,
																 $22, $23, $24)`,
		guest.Status, guest.FirstName, guest.LastName, guest.Gender,
		guest.UnitNum, guest.StAddress, guest.State, guest.City, guest.Zip, guest.TelNum, guest.Email,
		guest.ChildNum, guest.AdultNum, guest.PlaceOfWorship, guest.IsMember, guest.IsBaptized,
		guest.IsEspanol, guest.IsUnemployed, guest.IsHomeless, guest.IsFamily,
		guest.IsContactOk, guest.Allergies, guest.Notes)

	if err != nil {
		return err
	}

	result.RowsAffected()

	fmt.Println("Guest Successfully Saved.")

	return nil
}
