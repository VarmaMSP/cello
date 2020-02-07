import { Epic, ofType } from 'redux-observable'
import { from, Observable, of } from 'rxjs'
import {
  catchError,
  concatMap,
  debounceTime,
  filter,
  map,
  throttleTime,
} from 'rxjs/operators'
import { AppState } from 'store'
import * as T from 'types/actions'
import { UpdateTextAction } from 'types/actions/ui/search_bar'
import { doFetch } from 'utils/fetch'
import { qs } from 'utils/utils'

const searchEpic: Epic<T.AppActions, T.AppActions, AppState> = (action$) =>
  action$.pipe(
    ofType(T.SEARCH_BAR_UPDATE_TEXT),
    filter<UpdateTextAction>(({ text }) => text.trim().length > 0),
    debounceTime<UpdateTextAction>(350),
    concatMap<UpdateTextAction, Observable<T.AppActions>>((action) =>
      from(
        doFetch({
          method: 'POST',
          urlPath: `/ajax/service?${qs({
            endpoint: 'search_suggestions',
            query: action.text,
          })}`,
        }),
      ).pipe(
        map(({ podcastSearchResults }) => ({
          type: T.SEARCH_SUGGESTIONS_ADD_PODCAST,
          podcasts: podcastSearchResults,
        })),
        catchError<any, Observable<T.AppActions>>(() =>
          of({ type: 'CONTINUE' }),
        ),
      ),
    ),
    throttleTime<T.AppActions>(100),
  )

export default searchEpic
