export namespace api {
	
	export class CredentialsTest {
	    TogglTrackOk: boolean;
	    RedmineOk: boolean;
	
	    static createFrom(source: any = {}) {
	        return new CredentialsTest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.TogglTrackOk = source["TogglTrackOk"];
	        this.RedmineOk = source["RedmineOk"];
	    }
	}

}

export namespace cfg {
	
	export class ApiConfig {
	    key: string;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new ApiConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.url = source["url"];
	    }
	}
	export class Config {
	    toggl?: ApiConfig;
	    redmine?: ApiConfig;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.toggl = this.convertValues(source["toggl"], ApiConfig);
	        this.redmine = this.convertValues(source["redmine"], ApiConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace main {
	
	export class LoadedConfig {
	    Config: cfg.Config;
	    IsValid: boolean;
	
	    static createFrom(source: any = {}) {
	        return new LoadedConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Config = this.convertValues(source["Config"], cfg.Config);
	        this.IsValid = source["IsValid"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace toggltrack {
	
	export class AppTask {
	    Id: number;
	    Issue: number;
	    Comment: string;
	    Duration: number;
	    DecimalDuration: number;
	    PastDecimalDuration: number;
	    Sync: boolean;
	    // Go type: time
	    Date: any;
	    Description: string;
	    IsValid: boolean;
	    MatchedWithRedmine: boolean;
	
	    static createFrom(source: any = {}) {
	        return new AppTask(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Issue = source["Issue"];
	        this.Comment = source["Comment"];
	        this.Duration = source["Duration"];
	        this.DecimalDuration = source["DecimalDuration"];
	        this.PastDecimalDuration = source["PastDecimalDuration"];
	        this.Sync = source["Sync"];
	        this.Date = this.convertValues(source["Date"], null);
	        this.Description = source["Description"];
	        this.IsValid = source["IsValid"];
	        this.MatchedWithRedmine = source["MatchedWithRedmine"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class AskedTasks {
	    Entries: AppTask[];
	    HasRunningTask: boolean;
	
	    static createFrom(source: any = {}) {
	        return new AskedTasks(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Entries = this.convertValues(source["Entries"], AppTask);
	        this.HasRunningTask = source["HasRunningTask"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

