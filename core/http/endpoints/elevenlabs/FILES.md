# core/http/endpoints/elevenlabs/soundgeneration.go  
## Package: elevenlabs  
  
Imports:  
- github.com/gofiber/fiber/v2  
- github.com/mudler/LocalAI/core/backend  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/http/middleware  
- github.com/mudler/LocalAI/core/schema  
- github.com/mudler/LocalAI/pkg/model  
- github.com/rs/zerolog/log  
  
External data, input sources:  
- The code uses the ElevenLabs SoundGeneration endpoint to generate audio from the input text.  
- It relies on the input provided in the schema.ElevenLabsSoundGenerationRequest object.  
  
TODOs:  
- Support uploading files?  
  
### SoundGenerationEndpoint function  
This function handles the ElevenLabs SoundGeneration endpoint. It takes the input from the schema.ElevenLabsSoundGenerationRequest object and uses the backend.SoundGeneration function to generate audio. The generated audio is then downloaded by the client.  
  
### Summary  
The elevenlabs package provides a function to handle the ElevenLabs SoundGeneration endpoint. It takes the input text, duration, temperature, and other parameters to generate audio using the backend.SoundGeneration function. The generated audio is then downloaded by the client.  
  
# core/http/endpoints/elevenlabs/tts.go  
Package: elevenlabs  
  
Imports:  
- github.com/mudler/LocalAI/core/backend  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/http/middleware  
- github.com/mudler/LocalAI/pkg/model  
- github.com/gofiber/fiber/v2  
- github.com/mudler/LocalAI/core/schema  
- github.com/rs/zerolog/log  
  
External data, input sources:  
- The code uses the `schema.TTSRequest` struct to receive input parameters for the text-to-speech request.  
- It also relies on the `config.BackendConfigLoader`, `model.ModelLoader`, and `config.ApplicationConfig` to access backend configuration, model information, and application settings.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Summary:  
- The `TTSEndpoint` function handles the text-to-speech endpoint for the ElevenLabs TTS model. It receives the voice ID and the text to be converted to speech as input.  
- The function first checks if the input is valid and then uses the `backend.ModelTTS` function to generate the speech audio.  
- Finally, it returns the generated audio to the client.  
  
  
  
