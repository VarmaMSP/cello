import classnames from 'classnames'
import { iconMap } from 'components/icon'
import { SearchSuggestion } from 'models'
import React, { useEffect } from 'react'
import { getImageUrl, stopEventPropagation } from 'utils/dom'

export interface StateToProps {
  cursor: number
  suggestions: SearchSuggestion[]
}

export interface DispatchToProps {
  setCursor: (c: number) => void
  setSearchText: (text: string) => void
  loadResultsPage: (text: string) => void
  loadPodcastPage: (podcastUrlParam: string) => void
}

const SearchSuggestionsList: React.FC<StateToProps & DispatchToProps> = ({
  cursor,
  suggestions,
  setCursor,
  setSearchText,
  loadPodcastPage,
  loadResultsPage,
}) => {
  const handleOnClick = (c: number) => () => {
    setCursor(c)
    SearchSuggestion.isPodcast(suggestions[c])
      ? loadPodcastPage(suggestions[c].i)
      : loadResultsPage(suggestions[c].header.replace(/<\/?em>/g, ''))
  }

  const handleOnMouseOver = (c: number) => () => {
    setCursor(c)
  }

  const handleOnKeyDown = (e: any) => {
    switch (e.keyCode) {
      case 38:
        const c1 = cursor > 0 ? cursor - 1 : suggestions.length - 1
        setCursor(c1)
        if (SearchSuggestion.isTextSearch(suggestions[c1])) {
          setSearchText(suggestions[c1].header.replace(/<\/?em>/g, ''))
        }
        return

      case 40:
        const c2 = cursor < suggestions.length - 1 ? cursor + 1 : 0
        setCursor(c2)
        if (SearchSuggestion.isTextSearch(suggestions[c2])) {
          setSearchText(suggestions[c2].header.replace(/<\/?em>/g, ''))
        }
        return

      case 13:
        setCursor(cursor)
        SearchSuggestion.isPodcast(suggestions[cursor])
          ? loadPodcastPage(suggestions[cursor].i)
          : loadResultsPage(suggestions[cursor].header.replace(/<\/?em>/g, ''))
        return
    }
  }

  useEffect(() => {
    setCursor(0)
  }, [suggestions.length])

  useEffect(() => {
    document.addEventListener('keydown', handleOnKeyDown)
    return () => document.removeEventListener('keydown', handleOnKeyDown)
  }, [cursor, suggestions.map((x) => x.id).join()])

  return suggestions.length > 1 ? (
    <div
      style={{ width: '32rem' }}
      className="z-10 px-2 py-2 bg-white border border-blue-400 rounded-lg"
    >
      {suggestions
        .slice(1)
        .map((s, i) =>
          SearchSuggestion.isPodcast(s)
            ? renderItemSuggestion(
                s,
                handleOnClick(i + 1),
                handleOnMouseOver(i + 1),
                cursor === i + 1,
              )
            : renderTextSuggestion(
                s,
                handleOnClick(i + 1),
                handleOnMouseOver(i + 1),
                cursor === i + 1,
              ),
        )}
    </div>
  ) : (
    <></>
  )
}

function renderTextSuggestion(
  s: SearchSuggestion,
  onClick: () => void,
  onMouseOver: () => void,
  active: boolean,
): JSX.Element {
  const Icon = iconMap[s.i === 'S' ? 'search' : 'enter']

  return (
    <div
      key={s.id}
      className={classnames(
        'search-suggestion flex items-center px-3 py-1 cursor-pointer rounded',
        { 'bg-gray-200': active },
      )}
      onClick={onClick}
      onMouseOver={onMouseOver}
      onPointerDown={stopEventPropagation}
      onMouseDown={stopEventPropagation}
      onTouchStart={stopEventPropagation}
    >
      <Icon className="flex-none w-4 h-4 mr-4 fill-current text-gray-800" />
      <div
        className="lowercase text-base text-gray-900 leading-loose line-clamp-1"
        dangerouslySetInnerHTML={{ __html: s.header }}
      />
    </div>
  )
}

function renderItemSuggestion(
  s: SearchSuggestion,
  onClick: () => void,
  onMouseOver: () => void,
  active: boolean,
): JSX.Element {
  return (
    <div
      key={s.id}
      className={classnames(
        'search-suggestion flex p-3 hover:bg-gray-200 cursor-pointer rounded',
        { 'bg-gray-200': active },
      )}
      onClick={onClick}
      onMouseOver={onMouseOver}
      onPointerDown={stopEventPropagation}
      onMouseDown={stopEventPropagation}
      onTouchStart={stopEventPropagation}
    >
      <img
        src={getImageUrl(s.i)}
        className="flex-none w-12 h-12 mr-3 border rounded"
      />
      <div>
        <div
          className="text-base text-gray-900 line-clamp-1"
          dangerouslySetInnerHTML={{ __html: s.header }}
        />
        <div
          className="text-sm text-gray-900 line-clamp-1"
          dangerouslySetInnerHTML={{ __html: s.subHeader }}
        />
      </div>
    </div>
  )
}

export default SearchSuggestionsList
