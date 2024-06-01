package toggltrack

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"math"
	"sort"
	"time"
)

type ApiTask struct {
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Start       time.Time `json:"start"`
}

type AppTask struct {
	Issue               string
	Comment             string
	Duration            int
	DecimalDuration     float64
	PastDecimalDuration int
	Sync                bool
	Date                time.Time
	Description         string
	IsValid             bool
}

type AskedTasks struct {
	Entries        []*AppTask
	HasRunningTask bool
}

type appTasks map[string]*AppTask

func ProcessTasks(tasks []ApiTask) []*AppTask {
	e := groupTasks(tasks)
	computeDecimalDurations(&e)

	return sortTasks(e)
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
func computeDecimalDurations(e *appTasks) {
	for _, t := range *e {
		d := math.Round((float64(t.Duration)/3600)*4) / 4
		if d == 0 && t.Duration > 0 {
			t.DecimalDuration = .25
		} else {
			t.DecimalDuration = d
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
func sortTasks(e appTasks) []*AppTask {
	tasks := make([]*AppTask, 0, len(e))
	for _, t := range e {
		tasks = append(tasks, t)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Date.After(tasks[j].Date)
	})

	return tasks
}
