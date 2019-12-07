import ButtonPlay from 'components/button_play'
import EpisodeMeta from 'components/episode_meta'
import isToday from 'date-fns/isToday'
import isYesterday from 'date-fns/isYesterday'
import React from 'react'
import { Episode } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  history: Episode[]
  receivedAll: boolean
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadMore: (offset: number) => void
}

const Feed: React.FC<StateToProps & DispatchToProps> = ({ history }) => {
  const historyList: { title: string; episodes: Episode[] }[] = [
    { title: 'Today', episodes: [] },
    { title: 'Yesterday', episodes: [] },
    { title: 'Earlier', episodes: [] },
  ]

  for (let i = 0; i < history.length; ++i) {
    const episode = history[i]
    const pubDate = new Date(`${episode.pubDate} +0000`)

    if (isToday(pubDate)) {
      historyList[0].episodes.push(episode)
      continue
    }
    if (isYesterday(pubDate)) {
      historyList[1].episodes.push(episode)
      continue
    }
    historyList[2].episodes.push(episode)
  }

  return (
    <>
      {historyList.map(({ title, episodes }) => (
        <div key={title}>
          <h1 className="text-xl text-gray-900">{title}</h1>
          <hr className="mt-2 mb-6 border-gray-500" />
          {episodes.length > 0 ? (
            <div>
              {episodes.map((episode) => (
                <div key={episode.id} className="">
                  <img
                    className="w-24 h-24 flex-none object-contain rounded-lg border cursor-default"
                    src={getImageUrl(episode.podcastId)}
                    onClick={() => {}}
                  />
                  <div className="flex-auto flex flex-col justify-between pl-3">
                    <div>
                      <h1
                        className="text-sm leading-tight line-clamp-2 cursor-default"
                        onClick={() => {}}
                      >
                        {episode.title}
                      </h1>
                      <div className="mt-2">
                        <EpisodeMeta
                          displayPubDate={false}
                          episodeId={episode.id}
                        />
                      </div>
                    </div>
                    <ButtonPlay className="w-5" episodeId={episode.id} />
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <p className="my-6 text-gray-600 text-sm tracking-wide">
              {"Have'nt listenend to anything"}
            </p>
          )}
        </div>
      ))}
    </>
  )
}

export default Feed
