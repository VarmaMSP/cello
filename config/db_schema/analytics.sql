-- USERS
SELECT id, name, email, CONVERT_TZ(FROM_UNIXTIME(created_at), "+00:00", "+05:30") AS joined_at FROM user ORDER BY joined_at;
SELECT DATE(CONVERT_TZ(FROM_UNIXTIME(created_at), "+00:00", "+05:30")) AS joined_date, count(*) AS users_registered FROM user GROUP BY joined_date;

-- SUBSCRIPTIONS
SELECT user.name, COUNT(podcast_subscription.podcast_id) AS subscriptions FROM user INNER JOIN podcast_subscription ON podcast_subscription.subscribed_by = user.id GROUP BY user.name ORDER BY subscriptions;
SELECT podcast.title FROM podcast INNER JOIN podcast_subscription ON podcast_subscription.podcast_id = podcast.id WHERE podcast_subscription.subscribed_by = '';

-- EPISODE PLAYBACK
SELECT episode.title, user.name AS played_by, episode_playback.current_time_ AS played_til, CONVERT_TZ(episode_playback.last_played_at, "+00:00", "+05:30") AS played_at FROM episode_playback INNER JOIN episode ON episode.id = episode_playback.episode_id INNER JOIN user ON user.id = episode_playback.played_by WHERE user.id = 'bn1ajrmmvp8c9avf5190' ORDER BY played_at;
SELECT user.name, COUNT(episode_playback.episode_id) AS playbacks FROM user INNER JOIN episode_playback ON episode_playback.played_by = user.id GROUP BY user.name ORDER BY playbacks;

-- TASKS
SELECT name, type, `interval`, CONVERT_TZ(FROM_UNIXTIME(next_run_at), "+00:00", "+05:30") AS next_run_at, CONVERT_TZ(FROM_UNIXTIME(updated_at), "+00:00", "+05:30") AS updated_at from task; 