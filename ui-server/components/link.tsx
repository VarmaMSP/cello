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
        query: { podcastId: podcastId, activeTab: 'about' },
      }}
      as={`/podcasts/${podcastId}`}
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
