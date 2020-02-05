import NavTabs from 'components/nav_tabs'
import { NextSeo } from 'next-seo'
import React from 'react'
import { Episode } from 'types/models'
import EpisodeHeader from './components/episode_header'
import HomeTab from './tabs/home'

export interface OwnProps {
  episode: Episode
  activeTab: string | undefined
}

const EpisodeView: React.FC<OwnProps> = ({ episode, activeTab }) => {
  return (
    <div className="pt-6">
      <NextSeo
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
      <EpisodeHeader episode={episode} />
      <div className="mt-6 mb-4">
        <NavTabs
          tabs={[
            {
              name: 'episode',
              pathname: '/episodes',
              query: { episodeUrlParam: episode.urlParam, skipLoad: true },
              as: `/episodes/${episode.urlParam}`,
            },
          ]}
          active={activeTab}
          defaultTab="episode"
        />
      </div>
      <div className="mt-6 mb-4">
        <HomeTab episode={episode} />
      </div>
    </div>
  )
}

export default EpisodeView