package toggltrack

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"math"
	"sort"
	"time"

	"github.com/umanit/toggl-redmine/internal/redmine"
)

type ApiTask struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Start       time.Time `json:"start"`
}

type AppTask struct {
	Id                  int
	Issue               int
	Comment             string
	Duration            int
	DecimalDuration     float64
	PastDecimalDuration float64
	Sync                bool
	Date                time.Time
	Description         string
	IsValid             bool
	MatchedWithRedmine  bool
}

type AskedTasks struct {
	Entries        []*AppTask
	HasRunningTask bool
}

type appTasks map[string]*AppTask

// sameAsRedmineEntry vérifie si la tâche est équivalent à l’entrée de temps Redmine fournie.
func (task *AppTask) sameAsRedmineEntry(entry redmine.TimeEntry) bool {
	return entry.Issue.Id == task.Issue && entry.SpentOn == task.Date.Format(time.DateOnly) &&
		entry.Hours == task.DecimalDuration
}

// IsSyncable vérifie si la tâche est synchronisable.
func (task *AppTask) IsSyncable() bool {
	return !task.Sync || !task.IsValid || 0 == task.DecimalDuration
}

func ProcessTasks(tasks []ApiTask, timeEntries []redmine.TimeEntry) []*AppTask {
	t := groupTasks(tasks)
	computeDecimalDurations(&t)
	st := sortTasks(t)
	mutateWithRedmineEntries(st, timeEntries)

	return st
}

// UnmarshalJSON personnalisé uniquement pour corriger le fait que l’API de toggl track renvoie une durée négative si
// la tâche est encore en cours.
func (task *ApiTask) UnmarshalJSON(data []byte) error {
	type Alias ApiTask // Alias pour éviter la récursion infinie
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(task),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if task.Duration < 0 {
		task.Duration = 0
	}

	return nil
}

// groupTask regroupe les tâches ayant la même description et sur le même jour tout en vérifiant si leur description
// est valide.
func groupTasks(tasks []ApiTask) appTasks {
	e := make(appTasks)

	for _, t := range tasks {
		i := newIssue(t.Description)
		k := computeKey(t)

		if _, ok := e[k]; !ok {
			e[k] = &AppTask{
				Id:          t.Id,
				Issue:       i.Number,
				Comment:     i.Description,
				Sync:        i.IsValid,
				Date:        t.Start,
				Description: t.Description,
				IsValid:     i.IsValid,
			}
		}

		e[k].Duration += t.Duration
	}

	return e
}

// computeDecimalDuration convertie les durées des tâches au format décimal
func computeDecimalDurations(t *appTasks) {
	for _, task := range *t {
		d := math.Round((float64(task.Duration)/3600)*4) / 4
		if d == 0 && task.Duration > 0 {
			task.DecimalDuration = .25
		} else {
			task.DecimalDuration = d
		}
	}
}

// mutateWithRedmineEntries va vérifier si du temps a déjà été synchronisé sur Redmine
func mutateWithRedmineEntries(t []*AppTask, timeEntries []redmine.TimeEntry) {
	for _, entry := range timeEntries {
		for _, task := range t {
			if !task.MatchedWithRedmine && task.sameAsRedmineEntry(entry) {
				task.PastDecimalDuration = entry.Hours
				task.Sync = task.IsValid && 0 == entry.Hours
				task.MatchedWithRedmine = true
			}
		}
	}
}

// computeKey calcule la clef d’identification unique d’une tâche (même description sur un même jour)
func computeKey(t ApiTask) string {
	hash := sha1.New()
	hash.Write([]byte(t.Description + t.Start.Format(time.DateOnly)))
	hashBytes := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(hashBytes)
}

// sortTasks convertie la map en slice et trie les tâches par date décroissante
func sortTasks(t appTasks) []*AppTask {
	tasks := make([]*AppTask, 0, len(t))
	for _, task := range t {
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Date.After(tasks[j].Date)
	})

	return tasks
}
