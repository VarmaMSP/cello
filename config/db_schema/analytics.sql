-- USERS
SELECT id, name, email, CONVERT_TZ(FROM_UNIXTIME(created_at), "+00:00", "+05:30") AS joined_at FROM user ORDER BY joined_at;
SELECT DATE(CONVERT_TZ(FROM_UNIXTIME(created_at), "+00:00", "+05:30")) AS joined_date, count(*) AS users_registered FROM user GROUP BY joined_date;

-- SUBSCRIPTIONS
SELECT user.name, COUNT(subscription.podcast_id) AS subscriptions FROM user INNER JOIN subscription ON subscription.user_id = user.id GROUP BY user.name ORDER BY subscriptions;
SELECT podcast.title FROM podcast INNER JOIN subscription ON subscription.podcast_id = podcast.id WHERE subscription.user_id = '';

-- EPISODE PLAYBACK
SELECT episode.title, user.name AS user_id, playback.current_progress AS played_till, CONVERT_TZ(playback.last_played_at, "+00:00", "+05:30") AS played_at FROM playback INNER JOIN episode ON episode.id = playback.episode_id INNER JOIN user ON user.id = playback.user_id WHERE user.id = 5;
SELECT user.name, COUNT(playback.episode_id) AS playbacks FROM user INNER JOIN playback ON playback.user_id = user.id GROUP BY user.name ORDER BY playbacks;

-- TASKS
SELECT name, type, `interval`, CONVERT_TZ(FROM_UNIXTIME(next_run_at), "+00:00", "+05:30") AS next_run_at, CONVERT_TZ(FROM_UNIXTIME(updated_at), "+00:00", "+05:30") AS updated_at from task; 