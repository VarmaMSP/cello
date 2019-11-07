import NavTabs from 'components/nav_tabs'
import React from 'react'
import EpisodeList from './episode_list'
import PodcastInfo from './podcast_info'

interface OwnProps {
  podcastId: string
}

const PodcastView: React.FC<OwnProps> = ({ podcastId }) => {
  return (
    <div className="flex md:flex-row flex-col">
      <div className="lg:w-4/6 w-full">
        <PodcastInfo podcastId={podcastId} />
        <div className="mt-8 mb-4">
          <NavTabs tabs={['episodes', 'about']} active="episodes" />
        </div>
        <EpisodeList podcastId={podcastId} />
      </div>
    </div>
  )
}

export default PodcastView
