import ButtonShowMore from 'components/button_show_more'
import EpisodeListItem from 'components/episode_list_item'
import isToday from 'date-fns/isToday'
import isYesterday from 'date-fns/isYesterday'
import React from 'react'
import { Episode } from 'types/app'

export interface StateToProps {
  history: Episode[]
  receivedAll: boolean
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadMore: (offset: number) => void
}

const Feed: React.FC<StateToProps & DispatchToProps> = ({
  history,
  loadMore,
  receivedAll,
  isLoadingMore,
}) => {
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
                <EpisodeListItem key={episode.id} episodeId={episode.id} />
              ))}
            </div>
          ) : (
            <p className="my-6 text-gray-600 text-sm tracking-wide">
              {"Have'nt listenend to anything"}
            </p>
          )}
        </div>
      ))}

      {!receivedAll && (
        <div className="w-full h-10 mx-auto my-6">
          <ButtonShowMore
            isLoading={isLoadingMore}
            loadMore={() => loadMore(history.length)}
          />
        </div>
      )}
    </>
  )
}

export default Feed
