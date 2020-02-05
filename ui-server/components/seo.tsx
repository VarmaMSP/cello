import { NextSeo, NextSeoProps } from 'next-seo'
import React, { useEffect } from 'react'
import { Episode, Playlist, Podcast } from 'types/models'
import { getImageUrl } from 'utils/dom'
import { pageview } from 'utils/gtag'

const Seo: React.FC<NextSeoProps> = (props) => {
  useEffect(() => {
    pageview(window.location.pathname + window.location.search)
  }, [])

  return <NextSeo {...props} />
}

export const AboutPageSeo: React.FC<{}> = () => (
  <Seo
    title="About | Phenopod"
    description="Online Podcast Player"
    canonical="https://phenopod.com/about"
  />
)

export const ChartPageSeo: React.FC<{}> = () => (
  <Seo title="Chart | Phenopod" description="Chart" />
)

export const EpisodePageSeo: React.FC<{ episode: Episode }> = ({ episode }) => (
  <Seo
    title={`${episode.title} | Phenopod`}
    description={episode.summary}
    canonical={`https://phenopod.com/episodes/${episode.urlParam}`}
    openGraph={{
      url: `https://phenopod.com/episodes/${episode.urlParam}`,
      type: 'article',
      title: episode.title,
      description: episode.summary,
    }}
    twitter={{
      cardType: `summary_large_image`,
    }}
  />
)

export const ExplorePageSeo: React.FC<{}> = () => (
  <Seo
    title="Explore | Phenopod"
    description="Explore trending podcasts"
    canonical="https://phenopod.com/explore"
    openGraph={{
      url: 'https://phenopod.com/explore',
      type: 'website',
      title: 'Explore | Phenopod',
      description: 'Explore trending podcasts',
    }}
  />
)

export const HistoryPageSeo: React.FC<{}> = () => (
  <Seo
    noindex
    title="History | Phenopod"
    description="Browse your playback history"
    canonical="https://phenopod.com/history"
  />
)

export const IndexPageSeo: React.FC<{}> = () => (
  <Seo
    title="Phenopod"
    description="Online podcast player"
    canonical="https://phenopod.com"
  />
)

export const PlaylistPageSeo: React.FC<{ playlist: Playlist }> = ({
  playlist,
}) => (
  <Seo
    noindex={playlist.privacy === 'PRIVATE'}
    title={`${playlist.title} | Phenopod`}
    description={playlist.description}
    canonical={`https://phenopod.com/playlists/${playlist.urlParam}`}
    openGraph={{
      url: `https://phenopod.com/playlists/${playlist.urlParam}`,
      type: 'website',
      title: `${playlist.title} | Phenopod`,
      description: playlist.description,
    }}
  />
)

export const PlaylistsLibraryPageSeo: React.FC<{}> = () => (
  <Seo
    noindex
    title="Playlist Library | Phenopod"
    description="Your playlist library"
    canonical="https://phenopod.com/playlists"
  />
)

export const PodcastPageSeo: React.FC<{ podcast: Podcast }> = ({ podcast }) => (
  <Seo
    title={`${podcast.title} | Phenopod`}
    description={podcast.description}
    canonical={`https://phenopod.com/podcasts/${podcast.urlParam}`}
    openGraph={{
      url: `https://phenopod.com/podcasts/${podcast.urlParam}`,
      type: 'article',
      title: podcast.title,
      description: podcast.description,
      images: [{ url: getImageUrl(podcast.urlParam) }],
    }}
    twitter={{
      cardType: `summary_large_image`,
    }}
  />
)

export const PrivacyPageSeo: React.FC<{}> = () => (
  <Seo
    title="Privacy | Phenopod"
    description="Privacy"
    canonical="https://phenopod.com/privacy"
  />
)

export const ResultsPageSeo: React.FC<{ query: string }> = ({ query }) => (
  <Seo
    title={`${query} - Phenopod`}
    description={`${query} - Phenopod`}
    canonical={`https://phenopod.com/results?query=${query}`}
    openGraph={{
      url: `https://phenopod.com/results?query=${query}`,
      type: 'website',
      title: `${query} - Phenopod`,
      description: `Search results for ${query}`,
    }}
  />
)

export const SubscriptionsPageSeo: React.FC<{}> = () => (
  <Seo
    noindex
    title="Subscriptions | Phenopod"
    description="Browse latest episodes from your subscriptions"
    canonical="https://phenopod.com/subscriptions"
  />
)
