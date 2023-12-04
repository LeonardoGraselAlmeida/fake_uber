package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
)

type RideRepository struct {
	Db *sql.DB
}

type outputRideRepository struct {
	RideId      uuid.UUID
	PassengerId uuid.UUID
	DriverId    uuid.UUID
	Status      string
	Date        time.Time
	FromLat     float64
	FromLong    float64
	ToLat       float64
	ToLong      float64
}

func (repository *RideRepository) SaveRide(ride *entity.Ride) error {
	_, err := repository.Db.Exec("insert into cccat14.ride (ride_id, passenger_id, from_lat, from_long, to_lat, to_long, status, date) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		ride.RideId,
		ride.PassengerId,
		ride.FromLat,
		ride.FromLong,
		ride.ToLat,
		ride.ToLong,
		ride.GetStatus(),
		ride.Date)

	if err != nil {
		return err
	}

	return nil
}

func (repository *RideRepository) UpdateRide(ride *entity.Ride) error {
	_, err := repository.Db.Exec("update cccat14.ride set status = $1, driver_id = $2 where ride_id = $3",
		ride.GetStatus(), ride.GetDriverId(), ride.RideId)

	if err != nil {
		return err
	}

	return nil
}

func (repository *RideRepository) GetRideById(rideId uuid.UUID) (*entity.Ride, error) {
	var result outputRideRepository
	row := repository.Db.QueryRow("select ride_id, passenger_id, driver_id, status, from_lat, from_long, to_lat, to_long, date FROM cccat14.ride where ride_id = $1", rideId)

	if err := row.Scan(&result.RideId, &result.PassengerId, &result.DriverId, &result.Status, &result.FromLat, &result.FromLong, &result.ToLat, &result.ToLong, &result.Date); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	ride := entity.RestoreRide(result.RideId, result.PassengerId, result.DriverId, result.Status, result.Date, result.FromLat, result.FromLong, result.ToLat, result.ToLong)

	return ride, nil
}

func (repository *RideRepository) GetActiveRideByPassengerId(passengerId uuid.UUID) (*entity.Ride, error) {
	var result outputRideRepository
	row := repository.Db.QueryRow("select ride_id, passenger_id, driver_id, status, from_lat, from_long, to_lat, to_long, date FROM cccat14.ride where passenger_id = $1 and status in ('requested', 'accepted', 'in_progress')",
		passengerId)

	if err := row.Scan(&result.RideId, &result.PassengerId, &result.DriverId, &result.Status, &result.FromLat, &result.FromLong, &result.ToLat, &result.ToLong, &result.Date); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	ride := entity.RestoreRide(result.RideId, result.PassengerId, result.DriverId, result.Status, result.Date, result.FromLat, result.FromLong, result.ToLat, result.ToLong)

	return ride, nil
}

func (repository *RideRepository) GetAllRide() ([]*entity.Ride, error) {
	rows, err := repository.Db.Query("select  ride_id, passenger_id, driver_id, status, from_lat, from_long, to_lat, to_long, date from cccat14.ride")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rides []*entity.Ride
	for rows.Next() {
		var output outputRideRepository
		if err := rows.Scan(&output.RideId, &output.PassengerId, &output.DriverId, &output.Status, &output.FromLat, &output.FromLong, &output.ToLat, &output.ToLong, &output.Date); err != nil {
			log.Fatal("Erro ao ler os resultados: ", err)
			return nil, err
		}
		ride := entity.RestoreRide(output.RideId, output.PassengerId, output.DriverId, output.Status, output.Date, output.FromLat, output.FromLong, output.ToLat, output.ToLong)
		rides = append(rides, ride)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rides, nil
}
