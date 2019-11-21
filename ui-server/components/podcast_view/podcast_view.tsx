import NavTabs from 'components/nav_tabs'
import React from 'react'
import EpisodeList from './episode_list'
import PodcastAbout from './podcast_about'
import PodcastInfo from './podcast_info'

interface OwnProps {
  podcastId: string
  activeTab: string
}

const PodcastView: React.FC<OwnProps> = ({ podcastId, activeTab }) => {
  return (
    <div className="flex md:flex-row flex-col">
      <div className="lg:w-2/3 w-full">
        <PodcastInfo podcastId={podcastId} />
        <div className="mt-6 mb-4">
          <NavTabs
            tabs={[
              {
                name: 'episodes',
                pathname: '/podcasts',
                query: { podcastId, skipLoad: true, activeTab: 'episodes' },
                as: `/podcasts/${podcastId}/episodes`,
              },
              {
                name: 'about',
                pathname: '/podcasts',
                query: { podcastId, skipLoad: true, activeTab: 'about' },
                as: `/podcasts/${podcastId}/about`,
              },
            ]}
            active={activeTab}
          />
        </div>
        {activeTab === 'episodes' && <EpisodeList podcastId={podcastId} />}
        {activeTab === 'about' && <PodcastAbout podcastId={podcastId} />}
      </div>
    </div>
  )
}

export default PodcastView
