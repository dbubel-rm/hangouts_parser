package main

import (
	"encoding/json"
	"io/ioutil"
	"flag"
	"log"
	"sort"
	"fmt"
)

type HangoutJson struct {
	ContinuationEndTimestamp string `json:"continuation_end_timestamp"`
	ConversationState        []struct {
		ConversationID struct {
			ID string `json:"id"`
		} `json:"conversation_id"`
		ConversationState struct {
			ConversationID struct {
				ID string `json:"id"`
			} `json:"conversation_id"`
			Conversation struct {
				ID struct {
					ID string `json:"id"`
				} `json:"id"`
				Type                  string `json:"type"`
				SelfConversationState struct {
					SelfReadState struct {
						ParticipantID struct {
							GaiaID string `json:"gaia_id"`
							ChatID string `json:"chat_id"`
						} `json:"participant_id"`
						LatestReadTimestamp string `json:"latest_read_timestamp"`
					} `json:"self_read_state"`
					Status            string   `json:"status"`
					NotificationLevel string   `json:"notification_level"`
					View              []string `json:"view"`
					InviterID         struct {
						GaiaID string `json:"gaia_id"`
						ChatID string `json:"chat_id"`
					} `json:"inviter_id"`
					InviteTimestamp      string `json:"invite_timestamp"`
					SortTimestamp        string `json:"sort_timestamp"`
					ActiveTimestamp      string `json:"active_timestamp"`
					DeliveryMediumOption []struct {
						DeliveryMedium struct {
							MediumType string `json:"medium_type"`
						} `json:"delivery_medium"`
						CurrentDefault bool `json:"current_default"`
					} `json:"delivery_medium_option"`
					IsGuest bool `json:"is_guest"`
				} `json:"self_conversation_state"`
				ReadState []struct {
					ParticipantID struct {
						GaiaID string `json:"gaia_id"`
						ChatID string `json:"chat_id"`
					} `json:"participant_id"`
					LatestReadTimestamp string `json:"latest_read_timestamp"`
				} `json:"read_state"`
				HasActiveHangout   bool   `json:"has_active_hangout"`
				OtrStatus          string `json:"otr_status"`
				OtrToggle          string `json:"otr_toggle"`
				CurrentParticipant []struct {
					GaiaID string `json:"gaia_id"`
					ChatID string `json:"chat_id"`
				} `json:"current_participant"`
				ParticipantData []struct {
					ID struct {
						GaiaID string `json:"gaia_id"`
						ChatID string `json:"chat_id"`
					} `json:"id"`
					FallbackName        string `json:"fallback_name"`
					InvitationStatus    string `json:"invitation_status"`
					ParticipantType     string `json:"participant_type"`
					NewInvitationStatus string `json:"new_invitation_status"`
				} `json:"participant_data"`
				ForkOnExternalInvite   bool     `json:"fork_on_external_invite"`
				NetworkType            []string `json:"network_type"`
				ForceHistoryState      string   `json:"force_history_state"`
				GroupLinkSharingStatus string   `json:"group_link_sharing_status"`
			} `json:"conversation"`
			Event []struct {
				ConversationID struct {
					ID string `json:"id"`
				} `json:"conversation_id"`
				SenderID struct {
					GaiaID string `json:"gaia_id"`
					ChatID string `json:"chat_id"`
				} `json:"sender_id"`
				Timestamp      string `json:"timestamp"`
				SelfEventState struct {
					UserID struct {
						GaiaID string `json:"gaia_id"`
						ChatID string `json:"chat_id"`
					} `json:"user_id"`
					NotificationLevel string `json:"notification_level"`
				} `json:"self_event_state"`
				ChatMessage struct {
					MessageContent struct {
						Segment []struct {
							Type string `json:"type"`
							Text string `json:"text"`
						} `json:"segment"`
					} `json:"message_content"`
				} `json:"chat_message"`
				EventID               string `json:"event_id"`
				AdvancesSortTimestamp bool   `json:"advances_sort_timestamp"`
				EventOtr              string `json:"event_otr"`
				DeliveryMedium        struct {
					MediumType string `json:"medium_type"`
				} `json:"delivery_medium"`
				EventType    string `json:"event_type"`
				EventVersion string `json:"event_version"`
			} `json:"event"`
		} `json:"conversation_state"`
	} `json:"conversation_state"`
}

func (h *HangoutJson) GetAllChatText() {
	for _, convo := range h.ConversationState {
		participant_namess := []string{}
		for _, participant_names := range convo.ConversationState.Conversation.ParticipantData {
			participant_namess = append(participant_namess, participant_names.FallbackName)
		}
		sort.Slice(convo.ConversationState.Event, func(i, j int) bool { return convo.ConversationState.Event[i].Timestamp < convo.ConversationState.Event[j].Timestamp })
		for _, event := range convo.ConversationState.Event {
			for _,segment := range event.ChatMessage.MessageContent.Segment {
				fmt.Println(event.Timestamp, event.ConversationID.ID, participant_namess, segment.Text)
			}
		}
	}
}
func main() {
	var jsonfile string
	flag.StringVar(&jsonfile, "jsonfile","", "Google hangout JSON file to parse.")
	flag.Parse()

	if jsonfile == "" {
		log.Fatalln("Invalid params")
	}

	log.Println("Reading:", jsonfile)
	file, e := ioutil.ReadFile(jsonfile)
	if e != nil {
		log.Fatalln("Error reading .json file: %v\n", e)
	}
	var jsontype HangoutJson
	json.Unmarshal(file, &jsontype)
	jsontype.GetAllChatText()
}