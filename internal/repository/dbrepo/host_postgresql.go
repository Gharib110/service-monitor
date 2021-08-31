package dbrepo

import (
	"context"
	"database/sql"
	"github.com/DapperBlondie/service-monitor/internal/models"
	"github.com/rs/zerolog/log"
	"time"
)

// InsertHost uses for adding models.Host into Database
func (m *postgresDBRepo) InsertHost(h models.Host) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO hosts (host_name,canonical_name,url,ip,ipv6,location,os,active,created_at,updated_at)
               VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`
	var newID int
	err := m.DB.QueryRowContext(ctx, query,
		h.HostName,
		h.CanonicalName,
		h.URL,
		h.IP,
		h.IPV6,
		h.Location,
		h.OS,
		h.Active,
		h.CreatedAt,
		h.UpdatedAt,
	).Scan(&newID)
	if err != nil {
		log.Error().Msg(err.Error())
		return newID, err
	}

	return newID, nil
}

// GetHostByID uses for getting a specific models.Host by its ID
func (m *postgresDBRepo) GetHostByID(id int) (*models.Host, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT host_name,canonical_name,url,ip,ipv6,location,os,active,created_at,updated_at
				FROM hosts WHERE id=$1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var h *models.Host = &models.Host{}
	err := row.Scan(
		&h.HostName,
		&h.CanonicalName,
		&h.URL,
		&h.IP,
		&h.IPV6,
		&h.Location,
		&h.OS,
		&h.Active,
		&h.CreatedAt,
		&h.UpdatedAt,
		&h.ID,
	)
	if err != nil {
		log.Error().Msg(err.Error() + "; in getting values of query")
		return nil, err
	}

	return h, nil
}

// UpdateHost uses for updating a models.Host
func (m *postgresDBRepo) UpdateHost(h *models.Host) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `UPDATE hosts SET host_name=$1, canonical_name=$2, url=$3, ip=$4, ipv6=$5, os=$6,
		     active=$7, locaiton=$8, updated_at=$9 WHERE id=$10`
	_, err := m.DB.ExecContext(ctx, stmt,
		&h.HostName,
		&h.CanonicalName,
		&h.URL,
		&h.IP,
		&h.IPV6,
		&h.OS,
		&h.Active,
		&h.Location,
		time.Now(),
		&h.ID,
	)
	if err != nil {
		log.Error().Msg(err.Error() + "; in updating host by its ID")
		return err
	}

	return nil
}

func (m *postgresDBRepo) GetAllHosts() ([]*models.Host, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `SELECT id, host_name, canonical_name, url, ip, ipv6, location, os, active, created_at,
				updated_at FROM ORDER BY host_name`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Msg(err.Error())
			return
		}
	}(rows)
	var hosts []*models.Host = []*models.Host{}
	for rows.Next() {
		var h *models.Host = &models.Host{}
		err = rows.Scan(
			&h.ID,
			&h.HostName,
			&h.CanonicalName,
			&h.URL,
			&h.IP,
			&h.IPV6,
			&h.Location,
			&h.OS,
			&h.Active,
			&h.CreatedAt,
			&h.UpdatedAt,
		)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, err
		}
		hosts = append(hosts, h)
	}
	if err = rows.Err(); err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return hosts, nil
}
