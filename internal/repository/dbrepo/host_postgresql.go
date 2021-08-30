package dbrepo

import (
	"context"
	"github.com/DapperBlondie/service-monitor/internal/models"
	"github.com/rs/zerolog/log"
	"time"
)

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
