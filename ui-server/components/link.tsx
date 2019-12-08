import Link from 'next/link'
import React from 'react'

interface Props {
  children: JSX.Element
}

export const PodcastLink: React.FC<Props & { podcastId: string }> = ({
  children,
  podcastId,
}) => {
  return (
    <Link
      href={{
        pathname: '/podcasts',
        query: { podcastId: podcastId, activeTab: 'episodes' },
      }}
      as={`/podcasts/${podcastId}/episodes`}
      key={podcastId}
    >
      {children}
    </Link>
  )
}

export const EpisodeLink: React.FC<Props & { episodeId: string }> = ({
  children,
  episodeId,
}) => {
  return (
    <Link
      href={{
        pathname: '/episodes',
        query: { episodeId: episodeId, skipLoad: true },
      }}
      as={`/episodes/${episodeId}`}
    >
      {children}
    </Link>
  )
}

export const PodcastListLink: React.FC<Props & { listId: string }> = ({
  children,
  listId,
}) => {
  return (
    <Link
      href={{
        pathname: '/discover',
        query: { listId: listId },
      }}
      as={`/discover/${listId}`}
    >
      {children}
    </Link>
  )
}
