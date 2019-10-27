ALTER TABLE episode_playback ADD last_played_at DATETIME AFTER current_time_;
UPDATE episode_playback SET last_played_at = FROM_UNIXTIME(updated_at);
ALTER TABLE episode_playback MODIFY last_played_at DATETIME NOT NULL;