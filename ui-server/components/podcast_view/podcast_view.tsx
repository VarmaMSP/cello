import NavTabs from 'components/nav_tabs'
import React from 'react'
import { getIdFromUrlParam } from 'utils/format'
import EpisodeList from './episode_list'
import PodcastAbout from './podcast_about'
import PodcastInfo from './podcast_info'

interface OwnProps {
  podcastUrlParam: string
  activeTab: string
}

const PodcastView: React.FC<OwnProps> = ({ podcastUrlParam, activeTab }) => {
  const podcastId = getIdFromUrlParam(podcastUrlParam)

  return (
    <div className="flex md:flex-row flex-col">
      <div className="lg:w-2/3 w-full">
        <PodcastInfo podcastId={podcastId} />
        <div className="mt-6 mb-4">
          <NavTabs
            tabs={[
              {
                name: 'about',
                pathname: '/podcasts',
                query: { podcastUrlParam, skipLoad: true, activeTab: 'about' },
                as: `/podcasts/${podcastUrlParam}`,
              },
              {
                name: 'episodes',
                pathname: '/podcasts',
                query: { podcastUrlParam, skipLoad: true, activeTab: 'episodes' },
                as: `/podcasts/${podcastUrlParam}/episodes`,
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
