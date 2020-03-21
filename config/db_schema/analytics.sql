-- TASKS
SELECT name, type, interval_,
    CONVERT_TZ(FROM_UNIXTIME(next_run_at), "+00:00", "+05:30") AS next_run_at,
    CONVERT_TZ(FROM_UNIXTIME(updated_at), "+00:00", "+05:30") AS last_run_at from task;

-- USER ACTIVITY
SELECT id, name, email, CONVERT_TZ(FROM_UNIXTIME(created_at), "+00:00", "+05:30") AS joined_at FROM user
    ORDER BY joined_at;

-- NO OF SIGN UPS PER DAY
SELECT DATE(CONVERT_TZ(FROM_UNIXTIME(created_at), "+00:00", "+05:30")) AS joined_date, count(*) AS users_registered FROM user
    GROUP BY joined_date;

-- NO OF SUBSCRIPTIOS PER USER
SELECT user.name, COUNT(subscription.podcast_id) AS subscriptions FROM user
    INNER JOIN subscription ON subscription.user_id = user.id
    GROUP BY user.name
    ORDER BY subscriptions;

-- NO OF PLAYBACKS PER USER
SELECT user.name, COUNT(playback.episode_id) AS playbacks FROM user
    INNER JOIN playback ON playback.user_id = user.id
    GROUP BY user.name
    ORDER BY playbacks;

-- USER SUBSCRIPTIONS
SELECT user.name, podcast.title FROM subscription
    INNER JOIN user ON subscription.user_id = user.id
    INNER JOIN podcast ON subscription.podcast_id = podcast.id
    WHERE subscription.user_id = '';

-- USER PLAYBACKS
SELECT user.name, playback.current_progress, SEC_TO_TIME(episode.duration) AS duration, episode.title, DATE(CONVERT_TZ(playback.last_played_at, "+00:00", "+05:30")) AS played_at FROM playback
    INNER JOIN user ON playback.user_id = user.id
    INNER JOIN episode ON playback.episode_id = episode.id
    WHERE playback.user_id = '12'
    ORDER BY playback.last_played_at;

