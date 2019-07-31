CREATE OR REPLACE VIEW pantry.guest_visits_count_all
AS SELECT guest_visits.date_visit, count(guest_visits.date_visit) AS date_counts
   FROM pantry.guest_visits
  GROUP BY guest_visits.date_visit
  ORDER BY guest_visits.date_visit DESC;
