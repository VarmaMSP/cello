import ButtonShowMore from 'components/button_show_more'
import EpisodeListItem from 'components/episode_list_item'
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

const HistoryFeed: React.FC<StateToProps & DispatchToProps> = ({
  history,
  loadMore,
  receivedAll,
  isLoadingMore,
}) => {
  return (
    <div>
      <h1 className="text-xl text-gray-900">{`Your Listen History`}</h1>
      <hr className="mt-2 mb-6 border-gray-400" />
      {history.map((episode) => (
        <EpisodeListItem key={episode.id} episodeId={episode.id} />
      ))}

      {!receivedAll && (
        <div className="w-full h-10 mx-auto my-6">
          <ButtonShowMore
            isLoading={isLoadingMore}
            loadMore={() => loadMore(history.length)}
          />
        </div>
      )}
    </div>
  )
}

export default HistoryFeed
