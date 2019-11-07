import NavTabs from 'components/nav_tabs'
import React, { useState } from 'react'
import EpisodeList from './episode_list'
import PodcastAbout from './podcast_about'
import PodcastInfo from './podcast_info'

interface OwnProps {
  podcastId: string
  activeTab: string
}

const PodcastView: React.FC<OwnProps> = ({
  podcastId,
  activeTab: activeTab_,
}) => {
  const [activeTab, setActiveTab] = useState<string>(activeTab_)
  const handleTabClick = (tab: string) => tab !== activeTab && setActiveTab(tab)

  return (
    <div className="flex md:flex-row flex-col">
      <div className="lg:w-4/6 w-full">
        <PodcastInfo podcastId={podcastId} />
        <div className="mt-8 mb-4">
          <NavTabs
            tabs={['episodes', 'about']}
            active={activeTab}
            onClick={handleTabClick}
          />
        </div>
        {activeTab === 'episodes' && <EpisodeList podcastId={podcastId} />}
        {activeTab === 'about' && <PodcastAbout podcastId={podcastId} />}
      </div>
    </div>
  )
}

export default PodcastView
