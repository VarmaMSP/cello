import NavTabs from 'components/nav_tabs'
import React from 'react'
import { getIdFromUrlParam } from 'utils/format'
import EpisodeList from './episode_list'
import PodcastAbout from './podcast_about'
import PodcastInfo from './podcast_info'

interface OwnProps {
  podcastUrlParam: string
  activeTab?: string
}

const PodcastView: React.FC<OwnProps> = ({ podcastUrlParam, activeTab }) => {
  const podcastId = getIdFromUrlParam(podcastUrlParam)

  return (
    <div>
      <PodcastInfo podcastId={podcastId} />
      <div className="mt-6 mb-4">
        <NavTabs
          tabs={[
            {
              name: 'about',
              pathname: '/podcasts',
              query: { podcastUrlParam, skipLoad: true },
              as: `/podcasts/${podcastUrlParam}`,
            },
            {
              name: 'episodes',
              pathname: '/podcasts',
              query: {
                podcastUrlParam,
                skipLoad: true,
                activeTab: 'episodes',
              },
              as: `/podcasts/${podcastUrlParam}/episodes`,
            },
          ]}
          active={activeTab}
          defaultTab="about"
        />
      </div>
      {!!!activeTab && <PodcastAbout podcastId={podcastId} />}
      {activeTab === 'episodes' && <EpisodeList podcastId={podcastId} />}
    </div>
  )
}

export default PodcastView
