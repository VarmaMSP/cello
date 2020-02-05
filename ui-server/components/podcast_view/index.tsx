import NavTabs from 'components/nav_tabs'
import React from 'react'
import { Podcast } from 'types/models'
import PodcastHeader from './components/podcast_header/podcast_header'
import HomeTab from './tabs/home'

interface OwnProps {
  podcast: Podcast
  activeTab?: string
}

const PodcastView: React.FC<OwnProps> = ({ podcast, activeTab }) => {
  return (
    <div className="pt-6">
      <PodcastHeader podcast={podcast} />
      <div className="mt-6 mb-4">
        <NavTabs
          tabs={[
            {
              name: 'podcast',
              pathname: '/podcasts',
              query: { podcastUrlParam: podcast.urlParam, skipLoad: true },
              as: `/podcasts/${podcast.urlParam}`,
            },
          ]}
          active={activeTab}
          defaultTab="podcast"
        />
      </div>
      {activeTab === undefined && <HomeTab podcast={podcast} />}
    </div>
  )
}

export default PodcastView
