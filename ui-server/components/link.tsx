import Link from 'next/link'
import React from 'react'

interface Props {
  children: JSX.Element
}

export const PodcastLink: React.FC<Props & { podcastUrlParam: string }> = ({
  children,
  podcastUrlParam,
}) => {
  return (
    <Link
      href={{
        pathname: '/podcasts',
        query: { podcastUrlParam: podcastUrlParam },
      }}
      as={`/podcasts/${podcastUrlParam}`}
    >
      {children}
    </Link>
  )
}

export const EpisodeLink: React.FC<Props & { episodeUrlParam: string }> = ({
  children,
  episodeUrlParam,
}) => {
  return (
    <Link
      href={{
        pathname: '/episodes',
        query: { episodeUrlParam: episodeUrlParam },
      }}
      as={`/episodes/${episodeUrlParam}`}
    >
      {children}
    </Link>
  )
}

export const ChartLink: React.FC<Props & { chartId: string }> = ({
  children,
  chartId,
}) => {
  return (
    <Link
      href={{
        pathname: '/charts',
        query: { chartId: chartId },
      }}
      as={`/charts/${chartId}`}
    >
      {children}
    </Link>
  )
}

export const PlaylistLink: React.FC<Props & { playlistUrlParam: string }> = ({
  children,
  playlistUrlParam,
}) => {
  return (
    <Link
      href={{
        pathname: '/playlists',
        query: { playlistUrlParam: playlistUrlParam },
      }}
      as={`/playlists/${playlistUrlParam}`}
    >
      {children}
    </Link>
  )
}
